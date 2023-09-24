package router

import (
	"github.com/gofiber/fiber/v2"
)

func WebRouter(app *fiber.App) {

	// To render a template, you can call the ctx.Render function
	// Render(tmpl string, values interface{}, layout ...string)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Home Page",
			"Body":  "Lorem ipsum...",
		})
	})

	// Render with layout example
	app.Get("/layout", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

}
