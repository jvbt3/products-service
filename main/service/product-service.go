package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Aqui será criado os produtos!")
}

func ReadAllProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Aqui será listado todos os produtos!")
}

func UpdateProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Aqui será atualizado os produtos!")
}

func DeleteProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Aqui será deletado os produtos!")
}
