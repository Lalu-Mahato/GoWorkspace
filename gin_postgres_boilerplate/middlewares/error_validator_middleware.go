package middlewares

import (
	"fmt"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationMiddleware(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		validate := validator.New()

		if err := c.ShouldBindJSON(model); err != nil {
			handleValidationErrors(c, err)
			return
		}

		if err := validate.Struct(model); err != nil {
			handleValidationErrors(c, err)
			return
		}

		c.Set("validatedModel", model)
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

	utils.ErrorResponse(c, errorMessages)
	c.Abort()
}
