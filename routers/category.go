package routers

import (
	"encoding/json"
	"fmt"
	"gambit/db"
	"gambit/models"
	"strconv"
)
func validateBody(t models.Category, err error)(int, string){
	if err != nil {
		return 400, "Invalid Body" + err.Error()
	}
	if len(t.CategName) == 0{
		return 400, "Debe especificar el nombre de la categoria"
	}
	if len(t.CategPath) == 0{
		return 400, "Debe especificar el path de la categoria"
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