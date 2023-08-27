package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the application!",
		})
	})

	UserRoutes(router)
	AuthRoutes(router)
}
