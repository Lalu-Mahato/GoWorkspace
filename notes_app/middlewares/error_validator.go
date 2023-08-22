package middlewares

import (
	"fmt"
	"net/http"
	"notes_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validate := validator.New()
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			handleValidationErrors(c, err)
			return
		}

		if err := validate.Struct(user); err != nil {
			handleValidationErrors(c, err)
			return
		}

		c.Set("validatedUser", user)
		c.Next()
	}
}

func handleValidationErrors(c *gin.Context, err error) {
	errors := err.(validator.ValidationErrors)
	var errorMessages []string

	for _, e := range errors {
		message := fmt.Sprintf("Field: %s, Error: %s", e.Field(), e.Tag())
		errorMessages = append(errorMessages, message)
	}

	c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
	c.Abort()
}
