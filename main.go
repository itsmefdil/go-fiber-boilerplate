package main

import (
	"fiber/config/database"
	"fiber/router"
	"log"

	_ "fiber/docs/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	swagger "github.com/gofiber/swagger"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Connect to database
	database.ConnectDB()

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/health", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	router.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
