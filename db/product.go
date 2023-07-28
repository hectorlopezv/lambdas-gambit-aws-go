package db

import (
	"database/sql"
	"errors"
	"fmt"
	"gambit/models"
	"gambit/tools"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func InsertProduct(t models.Product)(int64, error){
	fmt.Println("Comienza registro de InsertProduct")
	err := DbConnect()
	if err != nil {
		return 400, err
	}
	defer Db.Close()
	query := "INSERT INTO products (Prod_Title "
	if len(t.ProdDescription) > 0{
		query += ", Prod_Description"
	}
	if t.ProdPrice > 0{
query += ", Prod_Price"
	}
	if t.ProdCategId > 0{
		query += ", Prod_CategoryId"
	}
	if t.ProdStock > 0{
		query += ", Prod_Stock"
	}
	if len(t.ProdPath) > 0{
		query += ", Prod_Path"
	}
	query += ") VALUES ('" + tools.EscapeString(t.ProdTitle) + "'"

	if len(t.ProdDescription) > 0{
		query += ", '" + tools.EscapeString(t.ProdDescription) + "'"
	}
	if t.ProdPrice > 0{
		query += ", " + strconv.FormatFloat(t.ProdPrice, 'e', -1, 64)
	}
	if t.ProdCategId > 0{
		query += ", " + strconv.Itoa(t.ProdCategId)
	}
	if t.ProdStock > 0{
		query += ", " + strconv.Itoa(t.ProdStock)
	}
	if len(t.ProdPath) > 0{
		query += ", '" + tools.EscapeString(t.ProdPath) + "'"
	}
	query += ")"
	fmt.Println(query)
	result, err := Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	lastInsertedId, err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println(err2.Error())
		return 0, err2
	}
	fmt.Println("Insert Product > Ejecuccion exitosa Insert Product")
	return lastInsertedId, nil
}
func UpdateProduct(t models.Product)error{
fmt.Println("comienza update")
err := DbConnect()
if err != nil {
	return err
}
defer Db.Close()
query := "UPDATE products SET "
query = tools.BuildQuery(query, "Prod_Title", "S", 0, 0, t.ProdTitle)
query = tools.BuildQuery(query, "Prod_Description", "S", 0, 0, t.ProdDescription)
query = tools.BuildQuery(query, "Prod_Price", "F", 0, t.ProdPrice, "")
query = tools.BuildQuery(query, "Prod_CategoryId", "N", t.ProdCategId, 0, "")
query = tools.BuildQuery(query, "Prod_Stock", "N", t.ProdStock, 0, "")
query = tools.BuildQuery(query, "Prod_Path", "S", 0, 0, t.ProdPath)
query += " WHERE Prod_Id = " + strconv.Itoa(t.ProdId)
fmt.Println(query)
_, err2 := Db.Exec(query)
if err2 != nil {
	fmt.Println(err2.Error())
	return err2
}
fmt.Println("Update Product > Ejecuccion exitosa Update Product")
return nil
}

func UpdateStock(t models.Product)error{
	fmt.Println("comienza update del stock")
	if t.ProdStock == 0 {
		return errors.New("Debe especificar el stock del producto")
	}

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()
	query := "UPDATE products SET Prod_Stock = Prod_Stock + " + strconv.Itoa(t.ProdStock) + " WHERE Prod_Id = " + strconv.Itoa(t.ProdId)
	fmt.Println(query)
	_, err2 := Db.Exec(query)
	if err2 != nil {
		fmt.Println(err2.Error())
		return err2
	}
	fmt.Println("Update Product > Ejecuccion exitosa Update Product")
	return nil
}
func SelectProduct(t models.Product, choice string, page int, pagesize int, ordertype string, orderfield string)(models.ProductResp, error){
	fmt.Println("Comienza registro de SelectProduct")
	var p models.ProductResp
	var prd []models.Product
	err := DbConnect()
	if err != nil {
		return p, err
	}
	defer Db.Close()
	var sentencia string
	var sentenciaCount string
	var where, limit string
	sentencia= "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Path, Prod_CategoryId, Prod_Stock FROM products "

	
	sentenciaCount= "SELECT count(*)  as registros FROM products "
	switch choice {
	case "P":
		where = " WHERE Prod_Id = " + strconv.Itoa(t.ProdId)
	case "C":
		where = " WHERE Prod_CategoryId = " + strconv.Itoa(t.ProdCategId)
	case "S":
		where = " WHERE UCASE(CONCAT(Prod_Title, Prod_Description)) LIKE '%'" + strings.ToUpper(t.ProdSearch) + "%'"
	case "U":
		where = " WHERE UCASE(Prod_Path) LIKE '%'"+ strings.ToUpper(t.ProdPath) + "%'"
	case "K":
		join:= "JOIN category ON Prod_CategoryId = Categ_Id AND Categ_Path LIKE '%'"+ strings.ToUpper(t.ProdCategPath) + "%' "
		sentencia += join
		sentenciaCount += join
	}
	sentenciaCount += where
	var rows *sql.Rows
	rows, err =Db.Query(sentenciaCount)

	if err != nil {
		fmt.Println(err.Error())
		return p, err
	}
	defer rows.Close()
rows.Next()
var regi sql.NullInt32
err = rows.Scan(&regi)
if err != nil {
	fmt.Println(err.Error())
	return p, err
}
registros := int(regi.Int32)

if page > 0 {
	if registros > pagesize{
		limit = " LIMIT " + strconv.Itoa(pagesize)
		if page > 1{
			offset := (page - 1) * pagesize
			limit += " OFFSET " + strconv.Itoa(offset)
		}
	} else{
		limit = ""
	}
}
var orderBy string
if len(orderfield)> 0{
	switch orderfield {
	case "I":
		orderBy = " ORDER BY Prod_Id "
	case "T":
		orderBy = " ORDER BY Prod_Title "
	case "D":
		orderBy = " ORDER BY Prod_Description "
	case "F":
		orderBy = " ORDER BY Prod_CreatedAt "
	case "P":
		orderBy = " ORDER BY Prod_Price "
	case "S":
		orderBy = " ORDER BY Prod_Stock "
	case "C":
		orderBy = " ORDER BY Prod_CategoryId "
	}
	if ordertype == "D"{
		orderBy += " DESC"
	} else{
		orderBy += " ASC"
	}
}
sentencia += where + orderBy + limit
fmt.Println(sentencia)
rows, err = Db.Query(sentencia)
if err != nil {
	fmt.Println(err.Error())
	return p, err
}
for rows.Next(){
	var pt models.Product
	var prodId sql.NullInt32
	var prodTitle sql.NullString
	var prodDescription sql.NullString
	var prodPrice sql.NullFloat64
	var prodPath sql.NullString
	var prodCategId sql.NullInt32
	var prodStock sql.NullInt32
	var prodCreatedAt sql.NullTime
	var prodUpdated sql.NullTime
	fmt.Println(rows.Columns())
	fmt.Println(rows.ColumnTypes())
	err := rows.Scan(&prodId, &prodTitle, &prodDescription, &prodCreatedAt, &prodUpdated, &prodPrice, &prodPath, &prodCategId, &prodStock)
	if err !=nil{
		fmt.Println(err.Error())
		return p, err
	}
	pt.ProdId = int(prodId.Int32)
	pt.ProdTitle = prodTitle.String
	pt.ProdDescription = prodDescription.String
	pt.ProdCreatedAt = prodCreatedAt.Time.String()
	pt.ProdUpdated = prodUpdated.Time.String()
	pt.ProdPrice = prodPrice.Float64
	pt.ProdPath = prodPath.String
	pt.ProdCategId = int(prodCategId.Int32)
	pt.ProdStock = int(prodStock.Int32)
	prd = append(prd, pt)
}

p.TotalItems = registros
p.Data = prd
fmt.Println("Select Product > Ejecuccion exitosa Select Product")
return p, nil

}
func DeleteProduct(id int)error{ 
	fmt.Println("Comienza registro de DeleteProduct")
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()
	sentencia := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(id)

	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Delete Category > Ejecuccion exitosa Delete category")
	return nil
}