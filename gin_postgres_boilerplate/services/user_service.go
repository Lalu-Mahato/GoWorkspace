package services

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/repositories"
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/utils"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService *UserService) GetUsers() ([]models.User, error) {
	users, err := userService.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService *UserService) CreateUser(user *models.User) error {
	hashedPassword := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	err := userService.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}
