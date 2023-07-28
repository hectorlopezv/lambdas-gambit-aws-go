package routers

import (
	"encoding/json"
	"fmt"
	"gambit/db"
	"gambit/models"
)

func InsertAddress(body string, user string)(int, string){
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		fmt.Println(err.Error())
		return 500, "Error al procesar la solicitud"
	}
	if t.AddAddress == "" {
		return 400, "Debe ingresar una direccion"
	}
	if t.AddCity == "" {
		return 400, "Debe ingresar una ciudad"
	}
	if t.AddState == "" {
		return 400, "Debe ingresar un estado"
	}
	if t.AddPostalCode == "" {
		return 400, "Debe ingresar un codigo postal"
	}
	if t.AddPhone == "" {
		return 400, "Debe ingresar un telefono"
	}
	if t.AddTitle == "" {
		return 400, "Debe ingresar un titulo"
	}
	err = db.InsertAddress(t, user)
	if err != nil {
		fmt.Println(err.Error())
		return 500, "Error al procesar la solicitud"
	}
	return 200, "Registro exitoso"
}
func UpdateAddress(body string, user string, id int)(int, string){
	var t models.Address
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		fmt.Println(err.Error())
		return 500, "Error al procesar la solicitud"
	}
	t.AddId = id
	var encontrado bool
	err, encontrado = db.AddressExists(user, t.AddId)
	if !encontrado {
	if err!= nil {
		return 500, "Error al procesar la solicitud"
	}
		return 404, "No se encontro el registro"
	}
	err = db.UpdateAddress(t)
	if err != nil {
		fmt.Println(err.Error())
		return 500, "Error al procesar la solicitud"
	}
	return 200, "Actualizacion exitosa"
	
}
func DeleteAddress(addressId int)(int,string){
	if addressId == 0 {
		return 400, "Debe especificar el id de la la direccion"
	}
	err2 := db.DeleteAddress(addressId)
	if err2 != nil {
		return 400, "Ocurrio un error al hacer delete a la adress"
	}
	return 200, "Delete Category > Ejecuccion exitosa Delete address"
}

func SelectAddress(user string)(int, string){
	add, err := db.SelectAddress(user)
	if err != nil {
		return 500, err.Error()
	}
	fmt.Println(add)
	resjson, err := json.Marshal(add)
	if err != nil {
		return 500, err.Error()
	}
	fmt.Println(string(resjson))
	return 200, string(resjson)
}