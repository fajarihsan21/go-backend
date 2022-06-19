package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func New() (*gorm.DB, error) {

	host := getEnv("HOST")
	user := getEnv("USER")
	password := getEnv("PASSWORD")
	dbName := getEnv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, user, password, dbName)

	gormDB, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, errors.New("database connection failed")
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, errors.New("database connection failed")
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return gormDB, nil
}
