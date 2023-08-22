package services

import (
	"notes_app/helpers"
	"notes_app/models"
	"notes_app/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (us *UserService) GetUsers() ([]models.User, error) {
	users, err := us.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) CreateUser(user *models.User) error {
	hashedPassword := helpers.HashPassword(user.Password)
	user.Password = hashedPassword
	err := us.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}
