package routers

import (
	"gin-restful-api/controllers"
	"github.com/gin-gonic/gin"
)	

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {
	   controllers.GetProducts(c)
	})
	r.GET("/product/:id", func(c *gin.Context) {
		controllers.GetProduct(c)
	})
	r.POST("/product", func(c *gin.Context) {
		controllers.CreateProduct(c);
	})
	r.PUT("/product/:id", func(c *gin.Context) {
		controllers.UpdateProduct(c);
	})
	r.DELETE("/product/:id", func(c *gin.Context) {
		controllers.DeleteProduct(c);
	})
	return r
}