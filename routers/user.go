package routers

import (
	"encoding/json"
	"fmt"
	"gambit/db"
	"gambit/models"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func UpdateUser(body string, user string)(int, string){
	var u models.User
	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		return 400, "Invalid Body" + err.Error()
	}
	if len(u.UserFirstName) == 0 && len(u.UserLastName) ==0{
		return 400, "Debe especificar el email del usuario (firstname) o (lastname)"
	}
	_, encontrado := db.UserExist(user)
	if !encontrado{
		fmt.Println("El usuario no existe" + user)
		return 400, "El usuario no existe"
	}

	err = db.UpdateUser(u, user)
	if err != nil {
		return 400, "Ocurrio un error al hacer update al usuario"
	}
	return 200, "Update User > Ejecuccion exitosa Update User"
}

func SelectUser(body string, user string)(int, string){
	fmt.Println("Select User")
	_, encontrado := db.UserExist(user)
	if !encontrado{
		fmt.Println("El usuario no existe" + user)
		return 400, "El usuario no existe"
	}
	row, err := db.SelectUser(user)
	if err != nil {
		return 400, "Ocurrio un error al hacer select al usuario"
	}
	fmt.Println(row)
	resjson, err := json.Marshal(row)
	if err != nil {
		return 400, "Ocurrio un error al hacer select al usuario"
	}

	fmt.Println(resjson)
	return 200, string(resjson)
}

func SelectUsers(body string, user string, request events.APIGatewayV2HTTPRequest)(int, string){
	var page int
	if len(request.QueryStringParameters["page"])==0{
		page = 1
	}else{
		page, _= strconv.Atoi(request.QueryStringParameters["page"])
	}
	isAdmin, msg := db.UserIsAdmin(user)
	if !isAdmin{
		return 400, msg
	}
	users, err := db.SelectUsers(page)
	if err != nil {
		return 400, "Ocurrio un error al hacer select al usuario"
	}
	resjson, err := json.Marshal(users)
	if err != nil {
		return 500, "Ocurrio un error al hacer select al usuario"
	}
	fmt.Println("Select Users exitoso")
	return 200, string(resjson)
}