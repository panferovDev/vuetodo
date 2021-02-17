package main

import (
	"github.com/panferovDev/vuetodo/back/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	models.ConnectDB()
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	route.Run(":3000")
}