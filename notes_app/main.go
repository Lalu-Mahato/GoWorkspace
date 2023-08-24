package main

import (
	"notes_app/config"
	"notes_app/routes"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initRouter() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()

	router.Use(helmet.Default())
	router.Use(cors.Default())

	config.ConnectDatabase()

	return router
}

func main() {
	r := initRouter()

	routes.SetupRouter(r)

	if err := r.Run(); err != nil {
		panic("Failed to start server")
	}
}
