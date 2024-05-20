package service

import (
	"example/products-service/main/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateProducts(context *gin.Context) {
	var newProduct model.ProductStruct

	if err := context.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	context.IndentedJSON(http.StatusCreated, newProduct)
}

func ReadAllProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, products)
}

func ReadOneProducts(context *gin.Context) {
	barcode := context.Param("barcode")

	convertedBarcode, err := strconv.ParseInt(barcode, 10, 64)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Erro": "Barcode não encontrado!"})
	}

	for _, a := range products {
		if a.Barcode == convertedBarcode {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func UpdateProducts(context *gin.Context) {
	var newProduct model.ProductStruct

	barcode := context.Param("barcode")

	convertedBarcode, err := strconv.ParseInt(barcode, 10, 64)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Erro": "Barcode não encontrado!"})
	}

	for i, a := range products {
		if a.Barcode == convertedBarcode {
			if err := context.BindJSON(&newProduct); err != nil {
				return
			}

			products[i] = newProduct
			context.IndentedJSON(http.StatusCreated, newProduct)
			return
		}
	}
}

func DeleteProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Aqui será deletado os produtos!")
}

var products = []model.ProductStruct{}
