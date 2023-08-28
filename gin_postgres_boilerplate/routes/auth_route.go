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

func AuthRoutes(r *gin.Engine) {
	db := config.DB

	userRepository := repositories.NewUserRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository, userRepository)
	authController := controllers.NewAuthController(authService)

	api := r.Group("/auth")
	api.POST("/login", middlewares.ValidationMiddleware(&models.Login{}), authController.LoginUser)
}
