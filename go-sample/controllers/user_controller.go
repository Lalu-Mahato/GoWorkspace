package controllers

import (
	"go-sample/models"
	"go-sample/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, _ := uc.UserService.FindUserByEmail(user.Email)
	if existingUser.Email != "" {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	if err := uc.UserService.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
