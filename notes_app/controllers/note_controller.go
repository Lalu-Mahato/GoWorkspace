package controllers

import (
	"net/http"
	"notes_app/models"
	"notes_app/services"
	"notes_app/utils"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	noteService *services.NoteService
}

func NewNoteController(noteService *services.NoteService) *NoteController {
	return &NoteController{noteService: noteService}
}

func (noteController *NoteController) CreateNote(c *gin.Context) {
	note := c.MustGet("validatedModel").(*models.Note)

	err := noteController.noteService.CreateNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}
	utils.CreatedResponse(c, note)
}

func (noteController *NoteController) FindNotes(c *gin.Context) {
	notes, err := noteController.noteService.FindNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
		return
	}
	c.JSON(http.StatusOK, notes)
}

// func FindNotes(c *gin.Context) {
// 	notes := []models.Note{}
// 	config.DB.Find(&notes)
// 	c.JSON(http.StatusOK, &notes)
// }

// func CreateNote(c *gin.Context) {
// 	var note models.Note
// 	c.BindJSON(&note)
// 	config.DB.Create(&note)
// 	c.JSON(http.StatusCreated, &note)
// }

// func DeleteNote(c *gin.Context) {
// 	var note models.Note
// 	config.DB.Where("id = ?", c.Param("id")).Delete(&note)
// 	c.JSON(http.StatusNoContent, nil)
// }

// func UpdateNote(c *gin.Context) {
// 	var note models.Note
// 	config.DB.Where("id = ?", c.Param("id")).First(&note)
// 	c.BindJSON(&note)
// 	config.DB.Save(&note)
// 	c.JSON(http.StatusOK, &note)
// }
