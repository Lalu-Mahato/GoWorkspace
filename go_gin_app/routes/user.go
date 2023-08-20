package routes

import (
	"go_test/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.FindUsers)
	router.POST("/", controller.CreateUser)
	router.PUT("/:id", controller.UpdateUser)
	router.DELETE("/:id", controller.DeleteUser)
}
