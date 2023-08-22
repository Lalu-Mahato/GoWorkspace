package routes

import (
	"notes_app/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(router *gin.Engine) {
	noteGroup := router.Group("/notes")
	{
		noteGroup.GET("/", controllers.FindNotes)
		noteGroup.POST("/", controllers.CreateNote)
		noteGroup.PUT("/:id", controllers.UpdateNote)
		noteGroup.DELETE("/:id", controllers.DeleteNote)
	}
}
