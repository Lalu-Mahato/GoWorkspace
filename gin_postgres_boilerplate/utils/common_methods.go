package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "Success",
		"data":   data,
	})
}

func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusCreated,
		"status": http.StatusText(http.StatusCreated),
		"data":   data,
	})
}

func DeleteResponse(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}

func NotFoundResponse(c *gin.Context, error interface{}) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":   http.StatusNotFound,
		"status": http.StatusText(http.StatusNotFound),
		"error":  error,
	})
}

func ErrorResponse(c *gin.Context, error interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":   http.StatusBadRequest,
		"status": http.StatusText(http.StatusBadRequest),
		"error":  error,
	})
}

func InternalServerErrorResponse(c *gin.Context, error interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":   http.StatusInternalServerError,
		"status": http.StatusText(http.StatusInternalServerError),
		"error":  error,
	})
}

func UnauthorizedResponse(c *gin.Context, error interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":   http.StatusUnauthorized,
		"status": http.StatusText(http.StatusUnauthorized),
		"error":  error,
	})
}
