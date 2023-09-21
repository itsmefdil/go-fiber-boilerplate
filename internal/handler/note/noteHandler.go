package noteHandler

import (
	"fiber/database"
	"fiber/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Notes Present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review Your Input", "data": err})
	}

	note.ID = uuid.New()

	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't Create Note", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note Created", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Note Found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note Found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Text     string `json:"text"`
	}

	db := database.DB
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Note Found", "data": nil})
	}

	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review Your Input", "data": err})
	}

	note.Title = updateNoteData.Title
	note.Subtitle = updateNoteData.Subtitle
	note.Text = updateNoteData.Text

	db.Save(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Note Updated", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Note Found", "data": nil})
	}

	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Couldn't Delete Note", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note Deleted"})
}
