package main

import (
	"go_test/config"
	"go_test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	config.Connect()
	routes.UserRoutes(r)
	r.Run()
}
