package controllers

import (
	"net/http"
	"strings"
	"task-manager-go/configs"
	"task-manager-go/inputs"
	"task-manager-go/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input inputs.CreateTaskInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		currentDateTime := time.Now()
		if !currentDateTime.Before(input.Deadline) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Deadline must be in future"})
			return
		}

		severities := map[string]bool{
			"low":    true,
			"medium": true,
			"high":   true,
		}

		if !severities[strings.ToLower(input.Severity)] {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid severity. Valid severities: low, medium, or high"})
			return
		}

		task := models.Task{
			Description: input.Description,
			Deadline:    input.Deadline,
			Severity:    input.Severity,
		}
		configs.DB.Create(&task)

		ctx.JSON(http.StatusCreated, gin.H{"task": task})
	}
}

func GetAllTasks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tasks []models.Task
		configs.DB.Find(&tasks)

		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
	}
}
