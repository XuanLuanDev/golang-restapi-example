package dals

import (
	"myapi/models"
	"myapi/utils"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func GetProducts() []models.Product {
	var products []models.Product

	db := GetConnection()
	defer db.Close()
	rows, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Code, &product.Desc, &product.Amount, &product.Status)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}

func CreateProduct(product *models.Product) bool {
	db := GetConnection()
	defer db.Close()
	stmt, err := db.Prepare("insert into products(name, price, code, description, amount, status) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(product.Name, product.Price, product.Code, product.Desc, product.Amount, product.Status)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if rowsAffected < 1 {
		return false
	}
	return true
}

func UpdateProduct(product *models.Product) bool {
	db := GetConnection()
	defer db.Close()
	stmt, err := db.Prepare("update products set name = ?, price = ?, code = ?, description = ?, amount = ?, status = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(product.Name, product.Price, product.Code, product.Desc, product.Amount, product.Status, product.ID)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if rowsAffected < 1 {
		return false
	}
	return true
}

func DeleteProduct(id string) bool {
	db := GetConnection()
	defer db.Close()
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if rowsAffected < 1 {
		return false
	}
	return true
}
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", utils.ConnectString)
	if err != nil {
		panic(err)
	}

	return db
}