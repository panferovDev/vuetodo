package main

import (
	"github.com/panferovDev/vuetodo/back/models"
	"github.com/panferovDev/vuetodo/back/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.Use(cors.Default())
	

	models.ConnectDB()

	route.GET("/", controllers.GetAllTasks)
	route.POST("/", controllers.CreateTask)
	route.DELETE("task/:id", controllers.DeleteTask)
	route.PUT("/task/:id", controllers.UpdateTask)

	route.Run(":3000")
}