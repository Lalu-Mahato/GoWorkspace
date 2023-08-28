package services

import (
	"errors"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/repositories"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/utils"
)

type AuthService struct {
	authRepository *repositories.AuthRepository
	userRepository *repositories.UserRepository
}

func NewAuthService(authRepository *repositories.AuthRepository, userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}

func (as *AuthService) LoginUser(email, password string) (*models.User, error) {
	user, err := as.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	isMatch := utils.CheckPasswordHash(password, user.Password)
	if !isMatch {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
