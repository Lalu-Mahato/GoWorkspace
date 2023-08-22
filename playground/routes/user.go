package routes

import (
	"playground/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/users", controllers.FindUsers)
	router.POST("/user", controllers.CreateUser)
}
