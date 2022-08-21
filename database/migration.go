package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/maha-1030/go-rest-api/models"
)

var db *gorm.DB
var err error

func init() {
	//Load .env file
	if err = godotenv.Load(); err != nil {
		fmt.Printf("Error occured while reading .env file, err: %v\n", err)

		return
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dialect := os.Getenv("DB_DIALECT")

	connectionURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbName, password)

	db, err = gorm.Open(dialect, connectionURL)
	if err != nil {
		fmt.Printf("Error occured while connecting to database, err: %v\n", err)
	}

	fmt.Printf("successfully connected to the %v database\n", dialect)

	db.Debug().AutoMigrate(&models.Customer{}, &models.Address{}, &models.Accounts{})
}

func GetDB() (*gorm.DB, error) {
	return db, err
}
