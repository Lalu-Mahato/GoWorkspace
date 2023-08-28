package main

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/config"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load the .env file")
	}
	config.DatabaseConnection()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	routes.SetupRouter(r)
	r.Run()
}
