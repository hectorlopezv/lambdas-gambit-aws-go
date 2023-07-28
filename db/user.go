package db

import (
	"database/sql"
	"fmt"
	"gambit/models"
	"gambit/tools"
	"strconv"

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

func SelectUser(userId string)(models.User, error){
	fmt.Println("Comienza registro de SelectUser")
	var u models.User
	err := DbConnect()
	if err != nil {
		return  u,err
	}
	defer Db.Close()
var rows *sql.Rows
	sentencia := "SELECT * FROM users WHERE User_UUID = '" + userId + "'"
	rows, err = Db.Query(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return u, err
	}
	defer rows.Close()
	rows.Next()
	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime
	rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateUpd, &dateUpg)
	u.UserFirstName = firstName.String
	u.UserLastName = lastName.String
	u.UserDateUpd = dateUpg.Time.String()
	fmt.Println("SelectUser > Ejecuccion exitosa SelectUser")
	return u, nil
}
func SelectUsers(page int)(models.ListUsers, error){
	fmt.Println("comeinza select users")
	var lu models.ListUsers
	user := []models.User{}
	err := DbConnect()
	if err != nil {
		return  lu,err
	}
	defer Db.Close()
	var offset int = (page - 1) * 10
	var sentencia string
	var sentenciaCount string = "SELECT count(*) as registros FROM users"
	sentencia ="SELECT * FROM users LIMIT 10"
	if offset > 0{
		sentencia += "OFFSET "+strconv.Itoa(offset)
	}
	var rowsCount *sql.Rows
	fmt.Println(sentenciaCount)
	rowsCount, err = Db.Query(sentenciaCount)
	if err != nil {
		fmt.Println("error en la consulta count")
		fmt.Println(err.Error())
		return lu, err
	}
	defer rowsCount.Close()
	rowsCount.Next()
	var registros int
	rowsCount.Scan(&registros)
	lu.TotalItems = registros

	
	var rows *sql.Rows
	fmt.Println(sentencia)
	rows, err = Db.Query(sentencia)
	if err != nil {
		fmt.Println("error en la consulta de sentencia")
		fmt.Println(err.Error())
		return lu, err
	}
	defer rows.Close()
	for rows.Next(){
		var u models.User
		var firstName sql.NullString
		var lastName sql.NullString
		var dateUpg sql.NullTime
		rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateUpd, &dateUpg)
		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpd = dateUpg.Time.String()
		user = append(user, u)
	}
lu.Data = user
	fmt.Println("SelectUser > Ejecuccion exitosa SelectUsers")
	return lu, nil

}