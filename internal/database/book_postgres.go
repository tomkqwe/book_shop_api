package database

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "postgres"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "postgres"

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(models.Book{})
	if err != nil {
		log.Fatal("Error connecting to the database... ", err)
	}
	fmt.Println("Database connection successful...")
}
