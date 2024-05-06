package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func InitDB() *gorm.DB {
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, port)

	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to the database")
	}

	return database
}

func DB() *gorm.DB {
	return database
}

func ApiPort() string {
	return os.Getenv("API_PORT")
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error closing database:", err)
		return
	}
	sqlDB.Close()
}
