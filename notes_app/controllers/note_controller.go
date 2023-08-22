package controllers

import (
	"net/http"
	"notes_app/config"
	"notes_app/models"

	"github.com/gin-gonic/gin"
)

func FindNotes(c *gin.Context) {
	notes := []models.Note{}
	config.DB.Find(&notes)
	c.JSON(http.StatusOK, &notes)
}

func CreateNote(c *gin.Context) {
	var note models.Note
	c.BindJSON(&note)
	config.DB.Create(&note)
	c.JSON(http.StatusCreated, &note)
}

func DeleteNote(c *gin.Context) {
	var note models.Note
	config.DB.Where("id = ?", c.Param("id")).Delete(&note)
	c.JSON(http.StatusNoContent, nil)
}

func UpdateNote(c *gin.Context) {
	var note models.Note
	config.DB.Where("id = ?", c.Param("id")).First(&note)
	c.BindJSON(&note)
	config.DB.Save(&note)
	c.JSON(http.StatusOK, &note)
}
