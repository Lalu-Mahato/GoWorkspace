package main

import (
	"fmt"
	"go-sample/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	environment := os.Getenv("ENV")
	fmt.Println("ENV: ", environment)

}

func main() {
	r := gin.Default()

	// Set up routes
	routes.SetupRoutes(r)

	r.Run()
}
