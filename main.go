package main

import (
	"fiber/config"
	"fiber/config/database"
	"fiber/router"
	"log"

	_ "fiber/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	swagger "github.com/gofiber/swagger"
)

// @title Fiber Swagger
// @version 2.0
// @description Golang with fiber.
// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	// Fiber instance
	app := fiber.New()

	// BasicAuth()
	// Provide a minimal config
	config := basicauth.Config{
		Users: map[string]string{
			config.Config("BASIC_AUTH_USERNAME"): config.Config("BASIC_AUTH_PASSWORD"),
			// Add more users as needed
		},
	}
	// Middleware for Basic Authentication
	app.Use(basicauth.New(config))

	// Connect to database
	database.ConnectDB()

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is up and running 🚀")
	})

	app.Get("/health", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	router.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health
// @Accept */health*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *fiber.Ctx) error {
	// Check DB connection
	res := map[string]interface{}{
		"status": "success",
		"database": map[string]interface{}{
			"status": "connected",
		},
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
