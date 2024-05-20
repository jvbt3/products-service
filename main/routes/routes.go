package routes

import (
	"example/products-service/main/service"
	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.POST("/products", service.CreateProducts)
	router.GET("/products", service.ReadAllProducts)
	router.GET("/products/:barcode", service.ReadOneProducts)
	router.PATCH("/products/:barcode", service.UpdateProducts)
	router.DELETE("/products", service.DeleteProducts)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
