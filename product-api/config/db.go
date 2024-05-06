package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	"product-res-api/model"

	"github.com/joho/godotenv"
)

var database *gorm.DB
var err error

// InitDB initializes the database connection 
// and performs auto migration for the Product model.
func InitDB() {
	// Load environment variables from.env file
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Database connection parameters
	host := os.Getenv("DB_HOST") //"localhost"
	user := os.Getenv("DB_USER") // "postgres"
	password := os.Getenv("DB_PASSWORD") // "postgres"
	dbName := os.Getenv("DB_NAME") // "db"
	port := os.Getenv("DB_PORT") // "5432"

	// Construct DSN (Data Source Name) for the PostgreSQL connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, port)

	// Open a connection to the PostgreSQL database
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to the database")
	}

	// Perform auto migration to create/update database tables based on the defined model structs
	database.AutoMigrate(&model.Product{})
	fmt.Println("AutoMigrate Product...")
	fmt.Println("Database connected")
}

// DB returns the global database instance.
func DB() *gorm.DB {
	return database
}

// ApiPort returns the port number to listen on.
func ApiPort() string {
	return os.Getenv("API_PORT")
}