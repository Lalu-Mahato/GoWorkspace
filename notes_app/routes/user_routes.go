package routes

import (
	"notes_app/config"
	"notes_app/controllers"
	"notes_app/middlewares"
	"notes_app/repositories"
	"notes_app/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	db := config.DB
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	api := router.Group("/users")
	api.GET("/", userController.FindUsers)
	api.POST("/", middlewares.ValidationMiddleware(), userController.CreateUser)
	api.POST("/upload", userController.UploadFile)
}
