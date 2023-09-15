package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"treasury-parser/config"
	"treasury-parser/models"
)

func Init() *gorm.DB {
	var host, user, password, dbname, port string = config.Config("POSTGRES_HOST"),
		config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"),
		config.Config("POSTGRES_DB"),
		config.Config("POSTGRES_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Entity{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
