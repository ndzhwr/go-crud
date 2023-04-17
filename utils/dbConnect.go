package utils

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)


var DB *gorm.DB ; 

func ConnectToDatabase() {
	var err error ;

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn) , &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
}