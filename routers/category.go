package routers

import (
	"encoding/json"
	"gambit/db"
	"gambit/models"
	"strconv"
)

// user from token
func InsertCategory(body string, user string)(int, string){
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Invalid Body" + err.Error()
	}
	if len(t.CategName) == 0{
		return 400, "Debe especificar el nombre de la categoria"
	}
	if len(t.CategPath) == 0{
		return 400, "Debe especificar el path de la categoria"
	}
	isAdmin, msg := db.UserIsAdmin(user)
	if !isAdmin{
		return 400, msg
	}
	result, err2 := db.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizsar el registro de la categoria "+ t.CategName+ " > " +err2.Error()
	}
	return 200, "{ categID: " + strconv.Itoa(int(result)) + "}"
}