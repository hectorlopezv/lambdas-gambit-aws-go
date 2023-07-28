package db

import (
	"fmt"
	"gambit/models"
	"strconv"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func AddressExists(user string, id int)(error, bool){
fmt.Println("Comienza registro de AddressExists")
err := DbConnect()
if err != nil {
	return  err, false
}
defer Db.Close()
sentencia := "SELECT 1 FROM addresses WHERE Add_Id = " + strconv.Itoa(id) + " AND Add_UserId = '" + user + "'"
fmt.Println(sentencia)
rows, err := Db.Query(sentencia)
if err != nil {
	fmt.Println(err.Error())
	return  err, false
}
defer rows.Close()
var valor string
rows.Next()
err = rows.Scan(&valor)
if err != nil {
	fmt.Println(err.Error())
	return  err, false
}

if valor == "1"{
	return nil, true
}
	return nil, false
}
func InsertAddress(addr models.Address, user string)error{
	fmt.Println("Comienza registro de InsertAddress")
	err := DbConnect()
	if err != nil {
		return  err
	}
	defer Db.Close()
	sentencia := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name)"
	sentencia += " VALUES ('"+user+"', '"+addr.AddAddress+"', '"+addr.AddCity+"', '"+addr.AddState+"', '"+addr.AddPostalCode+"', '"+addr.AddPhone+"', '"+addr.AddTitle+"', '"+addr.AddName+"')"
	_, err = Db.Exec(sentencia)
	if err != nil {
	fmt.Println(err.Error())
		return  err
	}
	fmt.Println(sentencia)
	fmt.Println("Finaliza registro de InsertAddress")
	return nil
}

func UpdateAddress(a models.Address)error{
fmt.Println("Update Address")
err := DbConnect()
if err != nil {
	return  err
}
defer Db.Close()
sentencia := "UPDATE addresses SET "
if a.AddAddress != ""{
	sentencia += "Add_Address = '" + a.AddAddress + "',"
}
if a.AddCity != ""{
	sentencia += "Add_City = '" + a.AddCity + "',"
}
if a.AddState != ""{
	sentencia += "Add_State = '" + a.AddState + "',"
}
if a.AddPostalCode != ""{
	sentencia += "Add_PostalCode = '" + a.AddPostalCode + "',"
}
if a.AddPhone != ""{
	sentencia += "Add_Phone = '" + a.AddPhone + "',"
}
if a.AddTitle != ""{
	sentencia += "Add_Title = '" + a.AddTitle + "',"
}
if a.AddName != ""{
	sentencia += "Add_Name = '" + a.AddName + "',"
}
sentencia, _= strings.CutSuffix(sentencia, ",")
sentencia += " WHERE Add_Id = " + strconv.Itoa(a.AddId)
fmt.Println(sentencia)
_, err = Db.Exec(sentencia)
	if err != nil {
	fmt.Println(err.Error())
		return  err
	}
	fmt.Println(sentencia)
	fmt.Println("Finaliza registro de UpdateAddress")
	return nil
}
func DeleteAddress(id int)error{
fmt.Println("Delete Address")
err := DbConnect()
if err != nil {
	return  err
}
defer Db.Close()
sentencia := "DELETE FROM addresses WHERE Add_Id = " + strconv.Itoa(id)
fmt.Println(sentencia)
_, err = Db.Exec(sentencia)

	if err != nil {
	fmt.Println(err.Error())
		return  err
	}
	
	fmt.Println(sentencia)
	fmt.Println("Finaliza registro de DeleteAddress")
	return nil
}
func SelectAddress(user string)([]models.Address, error){
	fmt.Println("Comienza registro de SelectAddress")
var t []models.Address
err := DbConnect()
if err != nil {
	return  t, err
}
defer Db.Close()
sentencia := "SELECT Add_Id, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name FROM addresses WHERE Add_UserId = '" + user + "'"
var rows *sql.Rows
rows, err = Db.Query(sentencia)
if err != nil {
	fmt.Println(err.Error())
	return  t, err
}
defer rows.Close()
for rows.Next() {
	var a models.Address
	var addId sql.NullInt16
	var addAddress sql.NullString
	var addCity sql.NullString
	var addState sql.NullString
	var addPostalCode sql.NullString
	var addPhone sql.NullString
	var addTitle sql.NullString
	var addName sql.NullString
	err = rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)
	if err != nil {
		fmt.Println(err.Error())
		return  t, err
	}
	a.AddId = int(addId.Int16)
	a.AddAddress = addAddress.String
	a.AddCity = addCity.String
	a.AddState = addState.String
	a.AddPostalCode = addPostalCode.String
	a.AddPhone = addPhone.String
	a.AddTitle = addTitle.String
	a.AddName = addName.String
	t = append(t, a)
}
fmt.Println("Finaliza registro de SelectAddress")

return t, nil
}