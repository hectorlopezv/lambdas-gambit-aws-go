package db

import (
	"database/sql"
	"fmt"
	"gambit/models"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateCategory( c models.Category)( error){
	fmt.Println("Comienza registro de UpdateCategory")
	err := DbConnect()
	if err != nil {
		return err
	}
	sentencia :=fmt.Sprintf("UPDATE category SET Categ_Name = '%s', Categ_Path = '%s' WHERE Categ_ID = %d", c.CategName, c.CategPath, c.CategID)
	defer Db.Close()

	result , err := Db.Exec(sentencia)
	if err != nil {
		return  err
	}
	

}
func InsertCategory( c models.Category)(int64, error){
	fmt.Println("Comienza registro de InsertCategory")
	 err := DbConnect()

	 if err != nil {
		 return 0, err
	 }
	 defer Db.Close()
	 sentencia := "INSERT INTO category (Categ_Name, Categ_Path) Values ('"+c.CategName+"', '"+c.CategPath+"')"
	 fmt.Println(sentencia)
	 var result sql.Result
	 result, err = Db.Exec(sentencia)
	 if err != nil {
	fmt.Println(err.Error())
		return 0, err
	 }
	 LastInsertId, err2 := result.LastInsertId()
	 if err2 != nil {
		fmt.Println(err2.Error())
	return 0, err2 
	}
	fmt.Println("Insert Category > Ejecuccion exitosa Insert category")
	return LastInsertId, nil

}