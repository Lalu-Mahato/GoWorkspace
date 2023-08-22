package main

import (
	"notes_app/config"
	"notes_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	config.ConnectDatabase()

	routes.SetupRouter(r)

	r.Run()
}
