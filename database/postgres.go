package database

import (
	"fmt"
	"log"

	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Book struct {
	Title         string          `json:"title"`
	Author        string          `json:"author"`
	YearPublushed uint            `json:"year_published"`
	Price         decimal.Decimal `json:"price"`
	Category      string          `json:"category"`
	gorm.Model
}

func DatabaseConnection() {
	host := "localhost"
	port := "5433"
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
	DB.AutoMigrate(Book{})
	if err != nil {
		log.Fatal("Error connecting to the database... ", err)
	}
	fmt.Println("Database connection successful...")
}
