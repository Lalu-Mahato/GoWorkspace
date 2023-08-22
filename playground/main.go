package main

import (
	"playground/config"
	"playground/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	config.Connect()

	routes.Routes(r)
	r.Run()
}
