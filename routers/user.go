package routers

import (
	"encoding/json"
	"fmt"
	"gambit/db"
	"gambit/models"
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