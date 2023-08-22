package controllers

import (
	"net/http"
	"notes_app/models"
	"notes_app/services"
	"notes_app/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) FindUsers(c *gin.Context) {
	users, err := uc.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	user := c.MustGet("validatedModel").(*models.User)

	err := uc.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	utils.CreatedResponse(c, user)
}

func (uc *UserController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile image upload failed"})
		return
	}

	utils.SuccessResponse(c, file)
}
