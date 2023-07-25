package db

import (
	"database/sql"
	"fmt"
	"gambit/models"
	"gambit/secretm"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB
func ReadSecret()error{
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
func DbConnect() error {
	Db, err = sql.Open("mysql",ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Conexion exitosa db")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string{
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = os.Getenv("DbName")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string)(bool, string){
	fmt.Println("UserIsAdmin start");

	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}
	defer Db.Close()

	query := "SELECT 1 FROM users WHERE User_UUID='"+userUUID+"' AND User_Status = 0'"
	fmt.Println(query)
	rows, err := Db.Query(query)
	if err != nil {
		return false, err.Error()
	}
	var value string
	rows.Next()
	rows.Scan(&value)
	fmt.Println("UserIsAdmin > Ejecucion exitosa - valor devuelto "+value);
	if value == "1"{
		return true, ""
	}
	return false, "User is not admin"
}