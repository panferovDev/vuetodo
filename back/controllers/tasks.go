package controllers

import (
	"github.com/panferovDev/vuetodo/back/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type CreateTaskData struct {
	Title string `json:"title" binding:"required"`
	Done  bool `json:"done"`
}

type UpdateData struct {
	Done bool 
}



func GetAllTasks(context *gin.Context) {

	var tasks []models.Task
	models.DB.Order("id").Find(&tasks)
	context.JSON(http.StatusOK, gin.H{"all": tasks})
}


func CreateTask(context *gin.Context) {
	var input CreateTaskData
	var tasks []models.Task
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := models.Task{Title: input.Title, Done: input.Done}
	models.DB.Create(&task)
	models.DB.Order("id").Find(&tasks)
	context.JSON(http.StatusOK, gin.H{"all": tasks})
}

func UpdateTask(context *gin.Context) {
	var task models.Task
	var tasks []models.Task
	if err := models.DB.Where("id = ?", context.Param("id")).First(&task).Error; err != nil {
		fmt.Println("ошибко")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateData
	if err := context.ShouldBindJSON(&input); err != nil {
		fmt.Println("ошибко dont exist")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&task).Update("done", input.Done)
	models.DB.Order("id").Find(&tasks)
	context.JSON(http.StatusOK, gin.H{"all": tasks})
}

func DeleteTask(context *gin.Context) {
	var task models.Task
	var tasks []models.Task
	if err := models.DB.Where("id = ?", context.Param("id")).First(&task).Error; err != nil {
		fmt.Println("ошибко")
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&task)
	models.DB.Order("id").Find(&tasks)
	context.JSON(http.StatusOK, gin.H{"all": tasks})
}