package controllers

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/constants"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/services"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) LoginUser(c *gin.Context) {
	loginData := c.MustGet("validatedModel").(*models.Login)

	user, err := ac.authService.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		utils.UnauthorizedResponse(c, constants.ErrorCodes["EA001"])
		return
	}

	token, error := utils.GenerateJWT(user.Email, user.ID)
	if error != nil {
		utils.ErrorResponse(c, constants.ErrorCodes["EA002"])
		return
	}
	loggedInData := models.LoggedInResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	utils.SuccessResponse(c, loggedInData)
}
