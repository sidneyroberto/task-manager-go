package main

import (
	"net/http"
	"task-manager-go/configs"
	"task-manager-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectToDatabase()

	routes.TaskRoute(router)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Task Manager API with Go"})
	})

	router.Run("localhost:8081")
}
