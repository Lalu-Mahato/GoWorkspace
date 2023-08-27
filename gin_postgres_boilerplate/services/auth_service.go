package services

import (
	"fmt"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/repositories"
)

type AuthService struct {
	authRepository *repositories.AuthRepository
}

func NewAuthService(authRepository *repositories.AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (authService *AuthService) LoginUser(login *models.Login) error {
	var user, err = authService.authRepository.GetUserByEmail(login.Email)
	if err != nil {
		return err
	}
	fmt.Println("user: ", user)
	return nil
}
