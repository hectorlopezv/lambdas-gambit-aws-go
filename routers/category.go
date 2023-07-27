package routers

import (
	"encoding/json"
	"fmt"
	"gambit/db"
	"gambit/models"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)
func validateBody(t models.Category, err error)(int, string){
	if err != nil {
		return 400, "Invalid Body" + err.Error()
	}
	if len(t.CategName) == 0 && len(t.CategPath) == 0{
		return 400, "Debe especificar el nombre de la categoria o path de la categoria"
	}

	return 200, ""
}
func validateisAdmin(user string)(bool, string){
	fmt.Println("user > "+user)
	isAdmin, msg := db.UserIsAdmin(user)
	if !isAdmin{
		return false, msg
	}
	return true, ""
}
// user from token
func InsertCategory(body string, user string)(int, string){
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)
	validation, vMessage  := validateBody(t, err)
	if validation != 200{
		return validation, vMessage
	}
isAdmin, msgA := validateisAdmin(user)
if !isAdmin{
	return 401, msgA
}


	result, err2 := db.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizsar el registro de la categoria "+ t.CategName+ " > " +err2.Error()
	}
	return 200, "{ categID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string, user string, id int)(int, string){
	if id == 0 {
		return 400, "Debe especificar el id de la categoria"
	}
 var t models.Category
 err := json.Unmarshal([]byte(body), &t)
 validation, vMessage  := validateBody(t, err)
 if validation != 200{
	 return validation, vMessage
 }
isAdmin, msgA := validateisAdmin(user)
if !isAdmin{
 return 401, msgA
}
t.CategID = id
 err2 := db.UpdateCategory(t)
 if err2 != nil {
	fmt.Println(err2.Error())
	return 400, "Ocurrio un error al hacer updatre a la categoria"
 }
 return 200, "Update Category > Ejecuccion exitosa Update category"
}

func DeleteCategory(body string, user string, id int)(int, string){
	if id == 0 {
		return 400, "Debe especificar el id de la categoria"
	}
	isAdmin, msgA := validateisAdmin(user)
	if !isAdmin{
		return 401, msgA
	}
	err2 := db.DeleteCategory(id)
	if err2 != nil {
		return 400, "Ocurrio un error al hacer delete a la categoria"
	}
	return 200, "Delete Category > Ejecuccion exitosa Delete category"
}

func SelectCategories(body string, request events.APIGatewayV2HTTPRequest)(int , string){
	var err error
	var CategId int
	var Slug string
	if len(request.QueryStringParameters["categId"])> 0 {
		CategId, err = strconv.Atoi(request.QueryStringParameters["categId"])
		if err != nil {
			return 400, "Invalid categId"
		}
	}else{
		if len(request.QueryStringParameters["slug"])> 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	lista, err2 := db.SelectCategories(CategId, Slug)
	if err2 != nil {
		return 400, "Ocurrio un error al hacer select a la categoria"
	}
	Categ, err3 := json.Marshal(lista)
	if err3 != nil {
		return 400, "Ocurrio un error al hacer select a la categoria"
	}
	return 200, string(Categ)
}