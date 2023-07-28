package routers

import (
	"encoding/json"
	"fmt"
	"gambit/db"
	"gambit/models"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)
func InsertOrder(body string, user string)(int, string){
	var o models.Orders
	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		fmt.Println(err.Error())
		return 500, "Error al procesar la solicitud"
	}
	o.Order_UserUUID = user

	ok, msg := validateOrder(o)
	if !ok {
		return 400, msg
	}
	result, err := db.InsertOrder(o)
	if err != nil {
		fmt.Println(err.Error())
		return 500, "Error al procesar la solicitud"
	}
	return 200, "{ OrderID: "+ string(result)+" }"
}

func validateOrder(o models.Orders)(bool, string){
	if o.Order_Total == 0 {
		return false, "Debe ingresar un total"
	}
	if o.Order_UserUUID == "" {
		return false, "Debe ingresar un usuario"
	}
	count := 0
	for _,od := range o.OrdersDetails{
		if od.OD_ProdId == 0 {
			return false, "Debe ingresar un producto"
		}
		if od.OD_Quantity == 0 {
			return false, "Debe ingresar una cantidad"
		}
		count ++
	}
	if count == 0 {
		return false, "Debe ingresar al menos un item en la order"
	}

	return true, ""
}


func SelectOrders(user string, request events.APIGatewayV2HTTPRequest)(int, string){
	var err error
	var fechaDesde, fechaHasta string
	var orderId int
	var page int
	if len(request.QueryStringParameters["fechaDesde"])>0 {
		fechaDesde = request.QueryStringParameters["fechaDesde"]
	}
	if len(request.QueryStringParameters["fechaHasta"])>0 {
		fechaHasta = request.QueryStringParameters["fechaHasta"]
	}
	if len(request.QueryStringParameters["page"])>0 {
		page,_= strconv.Atoi(request.QueryStringParameters["page"])
	}
	if len(request.QueryStringParameters["orderId"])>0 {
		orderId,_= strconv.Atoi(request.QueryStringParameters["orderId"])
	}
	result, err2 := db.SelectOrders(user, fechaDesde, fechaHasta, page, orderId)
	if err2 != nil {
		fmt.Println(err.Error())
		return 500, "Error al procesar la solicitud"
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err.Error())
		return 500, "Error al procesar la solicitud"
	}
	return 200, string(resultJson)
}