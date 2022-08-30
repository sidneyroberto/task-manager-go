package main

import (
	"task-manager-go/configs"
	"task-manager-go/routes"
)

func main() {
	router := routes.SetupRouter()

	dsn := configs.EnvMongoURI("DSN", "")
	configs.ConnectToDatabase(dsn)

	router.Run(":8081")
}
