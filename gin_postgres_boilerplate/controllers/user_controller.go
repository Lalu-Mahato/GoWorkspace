package controllers

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/services"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/utils"
	"net/http"

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
	userWithoutPassword := models.UserWithoutPassword{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	utils.CreatedResponse(c, userWithoutPassword)
}

func (uc *UserController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile image upload failed"})
		return
	}

	utils.SuccessResponse(c, file)
}
