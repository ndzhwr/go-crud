package main

import (
	"go-crud/models"
	"go-crud/utils"
	"log"
)

func main() {
	utils.LoadEnv()
	utils.ConnectToDatabase()

	err := utils.DB.AutoMigrate(&models.Post{});
	if err != nil {
		log.Fatal("Error when migrating the database")
	}
}