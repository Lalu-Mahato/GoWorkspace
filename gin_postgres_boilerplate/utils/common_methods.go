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
		"status": "Created",
		"data":   data,
	})
}

func DeleteResponse(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}

func NotFoundResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":   http.StatusNotFound,
		"status": "Not_Found",
		"data":   data,
	})
}

func ErrorResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":   http.StatusBadRequest,
		"status": "Bad_Request",
		"data":   data,
	})
}
