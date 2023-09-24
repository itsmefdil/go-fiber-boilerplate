package helpers

import (
	"fiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsHelper(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Config("CORS_ALLOW_ORIGINS"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

}
