package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("error loading .env file")
	}

	return os.Getenv(key)
}

func GetVersion() string {
	// Read file version.txt
	content, err := os.ReadFile("version.txt")
	if err != nil {
		log.Fatalf("Error reading version.txt: %v", err)
	}

	// Convert []byte to string and print to screen
	version := string(content)

	return version
}
