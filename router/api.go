package router

import (
	"fiber/app/helpers"
	noteRoutes "fiber/app/routes/note"
	"fiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func ApiRouter(app *fiber.App) {

	// Health Check
	app.Get("/health", config.HealthCheck)

	// Basic Auth
	helpers.BasicAuth(app)

	// Create a new Fiber instance
	api := app.Group("/api", logger.New())

	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault) // default

	noteRoutes.SetupNoteRoutes(api)

}
