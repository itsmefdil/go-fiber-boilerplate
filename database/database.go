package database

import (
	"fmt"
	"log"
	"strconv"

	"fiber/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB variable
var DB *gorm.DB

// ConnectDB function
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error parsing port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connection Opened to Database")
}
