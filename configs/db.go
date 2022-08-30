package configs

import (
	"fmt"
	"task-manager-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(dsn string) {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("App connected to database")

	database.AutoMigrate(&models.Task{})

	DB = database
}
