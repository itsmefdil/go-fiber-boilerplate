package helpers

import (
	"fiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuth(app *fiber.App) {
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
}
