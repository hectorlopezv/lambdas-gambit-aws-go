package db

import (
	"database/sql"
	"fmt"
	"gambit/models"
	"gambit/tools"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateCategory( c models.Category)( error){
	fmt.Println("Comienza registro de UpdateCategory")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()
	sentencia := "UPDATE category SET"

	if len(c.CategName) > 0{
		sentencia += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}
	if len(c.CategPath) > 0{
		if !strings.HasSuffix(sentencia, "SET"){
			sentencia += ", "
		}
		sentencia += " Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}
	sentencia += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Update Category > Ejecuccion exitosa Update category")
	return nil
}
func DeleteCategory(id int)error{
	fmt.Println("Comienza registro de DeleteCategory")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()
	sentencia := "DELETE FROM category WHERE Categ_Id = " + strconv.Itoa(id)

	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Delete Category > Ejecuccion exitosa Delete category")
	return nil
}
func SelectCategories(categId int, slug string)([]models.Category, error){

	fmt.Println("Comienza registro de SelectCategories")
	var Categ []models.Category
	err := DbConnect()
	if err != nil {
		return Categ, err
	}
	defer Db.Close()
	setencia := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category "
	if categId > 0 {
		setencia += " WHERE Categ_Id = " + strconv.Itoa(categId)
	} else {
		if len(slug) > 0 {
			setencia += " WHERE Categ_Path LIKE '%" + slug + "%'"
		}
	}
	fmt.Println(setencia)
	var rows *sql.Rows
	rows, err = Db.Query(setencia)
	if err != nil {
		fmt.Println(err.Error())
		return Categ, err
	}
	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err2 := rows.Scan(&categId, &categName, &categPath)
		if err2 != nil {
			fmt.Println(err2.Error())
			return Categ, err2
		}
		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String
		Categ = append(Categ, c)
	}
	fmt.Println("Select Categories > Ejecuccion exitosa Select Categories")
	return Categ, nil
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