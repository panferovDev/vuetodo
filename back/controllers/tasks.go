package controllers

import (
	"./vuetodo/back/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateTrackInput struct {
	Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct {
	Done string `json:"done"`
	Title  string `json:"title"`
}

func GetAllTasks(context *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	context.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTask(context *gin.Context) {
	var track models.Task
	if err := models.DB.Where("id = ?", context.Param("id")).First(&task).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}