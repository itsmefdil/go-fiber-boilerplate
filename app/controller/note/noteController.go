package noteController

import (
	"fiber/app/model"
	"fiber/config/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// NotesAPI godoc
// @Summary List all notes.
// @Description get all notes.
// @Tags Notes
// @ID get-notes
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/note [get]
func Get(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Notes Present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

// NotesAPI godoc
// @Summary Create a note.
// @Description create a note.
// @Tags Notes
// @ID create-notes
// @Param Body body model.NoteJson true "Create new note."
// @Example body {"title": "Note Title", "subtitle": "Note Subtitle", "text": "Note Text"}
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/note [post]
func Create(c *fiber.Ctx) error {
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

// NotesAPI godoc
// @Summary Get note by id.
// @Description Get note by id.
// @Tags Notes
// @ID getById-notes
// @Param uuid path string true "Note ID"
// @Example Body 1
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/note/{uuid} [get]
func GetById(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Note Found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note Found", "data": note})
}

// NotesAPI godoc
// @Summary Update a note by id.
// @Description Update a note by id.
// @Tags Notes
// @ID update-notes
// @Param uuid path string true "Note ID"
// @Param Body body model.NoteJson true "Create new note."
// @Example body {"title": "Note Title", "subtitle": "Note Subtitle", "text": "Note Text"}
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/note/{uuid} [put]
func Update(c *fiber.Ctx) error {
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

// NotesAPI godoc
// @Summary Delete note by id.
// @Description Delete note by id.
// @Tags Notes
// @ID deleteById-notes
// @Param uuid path string true "Note ID"
// @Example Body 1
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /api/note/{uuid} [delete]
func Delete(c *fiber.Ctx) error {
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
