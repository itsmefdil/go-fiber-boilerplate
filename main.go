package main

import (
	"encoding/json"
	"fiber/app/helpers"
	"fiber/config"
	"fiber/config/database"
	"fiber/router"
	"fmt"
	"log"

	_ "fiber/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

// @title Fiber Swagger
// @version 2.0
// @description Golang with fiber.
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {

	// Create a html engine
	engine := html.New("./resources/views", ".html")

	// Fiber instance
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Views:       engine,
	})

	// Favicons
	app.Static("/favicon.ico", "./public/img/favicon.png")

	// Connect to database
	database.ConnectDB()

	// mIddlewares
	app.Use(recover.New())
	helpers.CorsHelper(app)

	// Setup router
	router.WebRouter(app)
	router.ApiRouter(app)

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
