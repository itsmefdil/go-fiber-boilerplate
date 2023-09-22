package noteRoutes

import (
	noteController "fiber/app/controller/note"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	note.Post("/", noteController.Create)

	note.Get("/", noteController.Get)

	note.Get("/:noteId", noteController.GetById)

	note.Put("/:noteId", noteController.Update)

	note.Delete("/:noteId", noteController.Delete)

}
