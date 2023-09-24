package main

import (
	"encoding/json"
	"fiber/config"
	"fiber/config/database"
	"fiber/router"
	"fmt"
	"log"

	_ "fiber/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Fiber Swagger
// @version 2.0
// @description Golang with fiber.
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	// Fiber instance
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Favicons
	app.Static("/favicon.ico", "./public/img/favicon.png")

	// BasicAuth()
	// Provide a minimal config
	auth := basicauth.Config{
		Users: map[string]string{
			config.Config("BASIC_AUTH_USERNAME"): config.Config("BASIC_AUTH_PASSWORD"),
			// Add more users as needed
		},
	}
	// Middleware for Basic Authentication
	app.Use(basicauth.New(auth))

	// Connect to database
	database.ConnectDB()

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Config("CORS_ALLOW_ORIGINS"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Setup routes
	router.SetupRoutes(app)

	// Show Version
	fmt.Println("Application Version:", config.GetVersion())

	port := config.Config("APP_PORT")
	if port == "" {
		port = "3000"
	}

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
