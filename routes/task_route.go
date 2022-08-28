package routes

import (
	"task-manager-go/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoute(router *gin.Engine) {
	router.POST("/tasks", controllers.CreateTask())
	router.GET("/tasks", controllers.GetAllTasks())
}
