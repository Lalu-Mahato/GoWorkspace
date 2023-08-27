package routes

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/config"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/controllers"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/middlewares"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/repositories"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	db := config.DB
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	api := router.Group("/users")
	api.GET("/", userController.FindUsers)
	api.POST("/", middlewares.ValidationMiddleware(&models.User{}), userController.CreateUser)

}
