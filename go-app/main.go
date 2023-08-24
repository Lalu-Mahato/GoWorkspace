package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18"`
}

type Product struct {
	Name  string `json:"name" validate:"required,string"`
	Price int    `json:"price" validate:"gte=0 int"`
}

func ValidatorMiddleware(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		validate := validator.New()
		if err := validate.Struct(model); err != nil {
			var validationErrors []string
			for _, e := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors, fmt.Sprintf("Field: %s, Error: %s", e.Field(), e.Tag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			c.Abort()
			return
		}

		c.Set("validatedModel", model)
		c.Next()
	}
}

func CreateUser(c *gin.Context) {
	user, _ := c.Get("validatedModel")

	c.JSON(http.StatusOK, gin.H{"message": "User Created", "user": user})
}

func CreateProduct(c *gin.Context) {
	product, _ := c.Get("validatedModel")

	c.JSON(http.StatusOK, gin.H{"message": "Product Created", "product": product})
}

func main() {
	r := gin.Default()

	var user User
	var product Product
	r.POST("/create-user", ValidatorMiddleware(&user), CreateUser)
	r.POST("/create-product", ValidatorMiddleware(&product), CreateProduct)

	r.Run(":8080")
}
