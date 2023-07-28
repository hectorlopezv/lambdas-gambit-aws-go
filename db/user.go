package db

import (
	"fmt"
	"gambit/models"
	"gambit/tools"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateUser(u models.User, user string)error {
	fmt.Println("Comienza registro de UpdateUser")
	err := DbConnect()
	if err != nil {
		return  err
	}
	defer Db.Close()
	sentencia := "UPDATE users SET "
	coma:=""
	if len(u.UserFirstName) > 0{
		coma = ","
		sentencia +=  "User_FirstName = '" + u.UserFirstName + "'"
	}
	if len(u.UserLastName) > 0{
		sentencia += coma + " User_LastName = '" + u.UserLastName + "'"
	}
	sentencia += ", User_DateUpg = '" + tools.FechaMysql() + "' WHERE User_UUID = '" + user + "'"
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
fmt.Println("UpdateUser > Ejecuccion exitosa UpdateUser")
	return nil
}