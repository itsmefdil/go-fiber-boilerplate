package router

import (
	noteRoutes "fiber/app/routes/note"
	"fiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is up and running 🚀")
	})

	// Health Check
	app.Get("/health", config.HealthCheck)

	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault) // default

	noteRoutes.SetupNoteRoutes(api)

}
