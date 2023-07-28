package routers

import (
	"encoding/json"
	"fmt"
	"gambit/db"
	"gambit/models"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func InsertProduct (body string, user string)(int, string){
	fmt.Println("InsertProduct")
	var t models.Product
	err := json.Unmarshal([]byte(body), &t)
	
	if err != nil {
		return 400, "Invalid Body" + err.Error()
	}
	if len(t.ProdTitle) == 0{
		return 400, "Debe especificar el titulo del producto"
	}
	isAdmin, msgA := validateisAdmin(user)
	if !isAdmin{
		return 401, msgA
	}
	result , err2 := db.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizsar el registro del producto "+ t.ProdTitle+ " > " +err2.Error()
	}
	return 200, "{ prodID: " + strconv.Itoa(int(result)) + "}"
}
func UpdateProduct(body string, user string, id int)(int, string){
	if id == 0 {
		return 400, "Debe especificar el id del producto"
	}
 var t models.Product
 err := json.Unmarshal([]byte(body), &t)
 if err != nil {
	 return 400, "Invalid Body" + err.Error()
 }

isAdmin, msgA := validateisAdmin(user)
if !isAdmin{
 return 401, msgA
}
t.ProdId = id
 err2 := db.UpdateProduct(t)
 if err2 != nil {
	fmt.Println(err2.Error())
	return 400, "Ocurrio un error al hacer updatre a la producto"
 }
 return 200, "Update Product > Ejecuccion exitosa Update Product"
}

func UpdateStockProduct(body string, user string, id int)(int, string){
	if id == 0 {
		return 400, "Debe especificar el id del producto"
	}
 var t models.Product
 err := json.Unmarshal([]byte(body), &t)
 if err != nil {
	 return 400, "Invalid Body" + err.Error()
 }

isAdmin, msgA := validateisAdmin(user)
if !isAdmin{
 return 401, msgA
}
t.ProdId = id
 err2 := db.UpdateStock(t)
 if err2 != nil {
	fmt.Println(err2.Error())
	return 400, "Ocurrio un error al hacer updatre a la stock del producto"
 }
 return 200, "Update Product > Ejecuccion exitosa Update Product stock"
}
func DeleteProduct(body string, user string, id int)(int, string){
	if id == 0 {
		return 400, "Debe especificar el id de la producto"
	}
	isAdmin, msgA := validateisAdmin(user)
	if !isAdmin{
		return 401, msgA
	}
	err2 := db.DeleteProduct(id)
	if err2 != nil {
		return 400, "Ocurrio un error al hacer delete a la producto"
	}
	return 200, "Delete Category > Ejecuccion exitosa Delete producto"
}
func SelectProduct (request events.APIGatewayV2HTTPRequest)(int, string){
	var t models.Product
	var page, pagesize int
	var orderType, orderField string
	
	param := request.QueryStringParameters

	page ,_ = strconv.Atoi(param["page"])
	pagesize ,_ = strconv.Atoi(param["pagesize"])
	orderType = param["ordertype"] // D = Desc, A = Ascending, default = A
	orderField = param["orderfield"] // I or null, T title, D description, F Created At
	// P price , C CategId, S Stock

	if !strings.Contains("ITDFPCS", orderField){
		orderField = ""
	}
	var choice string

	if len(param["prodId"])> 0 {
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}
	if len(param["search"])>0{
		choice = "S"
		t.ProdSearch = param["search"]
	}
	if len(param["categId"])>0{
		choice = "C"
		t.ProdCategId, _ = strconv.Atoi(param["categId"])
	}
	if len(param["slug"])>0{
		choice = "U"
		t.ProdPath = param["slug"]
	}
	if len(param["slugCateg"])>0{
		choice = "K"
		t.ProdCategPath = param["slugCateg"]
	}
fmt.Println(param)


result, err2 := db.SelectProduct(t, choice, page, pagesize, orderType, orderField)
if err2 != nil {
	fmt.Println(err2.Error())
	return 400, "Ocurrio un error al hacer select a la producto"
}
product, err3 := json.Marshal(result)
if err3 != nil {
	fmt.Println(err3.Error())
	return 400, "Ocurrio un error al hacer select a la producto"
}
return 200, string(product)
} 