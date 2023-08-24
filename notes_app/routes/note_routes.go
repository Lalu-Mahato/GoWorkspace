package routes

import (
	"notes_app/config"
	"notes_app/controllers"
	"notes_app/middlewares"
	"notes_app/models"
	"notes_app/repositories"
	"notes_app/services"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(router *gin.Engine) {
	db := config.DB
	noteRepository := repositories.NewNoteRepository(db)
	noteService := services.NewNoteService(noteRepository)
	noteController := controllers.NewNoteController(noteService)

	api := router.Group("/notes")
	api.POST("/", middlewares.ValidationMiddleware(&models.Note{}), noteController.CreateNote)
	api.GET("/", noteController.FindNotes)

	// noteGroup := router.Group("/notes")
	// {
	// 	// noteGroup.GET("/", controllers.FindNotes)
	// 	// noteGroup.PUT("/:id", controllers.UpdateNote)
	// 	// noteGroup.DELETE("/:id", controllers.DeleteNote)
	// }
}
