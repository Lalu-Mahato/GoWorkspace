package routes

import (
	"go-sample/config"
	"go-sample/controllers"
	"go-sample/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_, database := config.DBinstance()

	collectionName := "users"
	userService := services.NewUserService(database, collectionName)
	userController := controllers.NewUserController(userService)

	api := router.Group("/api")
	api.POST("/register", userController.RegisterUser)
}
