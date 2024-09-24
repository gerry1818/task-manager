package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required,oneof=todo in-progress done"`
	UserID      uint   `json:"user_id"`
}

var DB *gorm.DB

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDatabase() {
	// Load environment variables
	LoadEnv()

	// Fetch database credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbTimezone := os.Getenv("DB_TIMEZONE")

	// PostgreSQL connection string
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=" + dbTimezone

	// Connect to the database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto migrate the schema
	database.AutoMigrate(&User{}, &Task{})

	DB = database
	log.Println("Database connection successful!")
}
