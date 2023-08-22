package repositories

import (
	"notes_app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur *UserRepository) Create(user *models.User) error {
	err := ur.db.Create(user).Error
	return err
}
