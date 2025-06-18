package controllers

import (
	"gin-restful-api/dals"
	"gin-restful-api/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context){
	products,err := dals.GetProducts()
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, products)
	}
}
func GetProduct(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product, err := dals.GetProduct(int(id))
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, product)
	}
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err == nil {
		dals.CreateProduct(&product)
		c.JSON(200, product)
	}else {
		c.JSON(500, "The product data incorrect")
	}
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var product models.Product
	if err := c.ShouldBindJSON(&product); err == nil {
		product.ID = int(id)
		dals.UpdateProduct(&product)
		c.JSON(200, product)
	}else {
		c.JSON(500, "The product data incorrect")
	}
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	dals.DeleteProduct(&models.Product{ID: int(id)})
	c.JSON(200, "Delete product success")
}