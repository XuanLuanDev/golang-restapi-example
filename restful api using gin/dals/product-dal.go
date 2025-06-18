package dals

import (
	"gin-restful-api/utils"
	"gin-restful-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(utils.ConnectString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetProduct(id int) (*models.Product, error) {
	db, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var product models.Product
	err = db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func GetProducts() ([]models.Product, error) {
	db, err := GetConnection()
	if err != nil {
		return nil, err
	}
	var products []models.Product
	err = db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func CreateProduct(product *models.Product) error {
	db, err := GetConnection()
	if err != nil {
		return err
	}
	return db.Create(product).Error
}

func UpdateProduct(product *models.Product) error {
	db, err := GetConnection()
	if err != nil {
		return err
	}
	return db.Save(product).Error
}

func DeleteProduct(product *models.Product) error {
	db, err := GetConnection()
	if err != nil {
		return err
	}
	return db.Delete(product).Error
}