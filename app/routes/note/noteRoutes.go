package noteRoutes

import (
	noteController "fiber/app/controller/note"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	note.Post("/", noteController.CreateNotes)

	note.Get("/", noteController.GetNotes)

	note.Get("/:noteId", noteController.GetNote)

	note.Put("/:noteId", noteController.UpdateNote)

	note.Delete("/:noteId", noteController.DeleteNote)

}
