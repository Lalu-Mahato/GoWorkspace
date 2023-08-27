package repositories

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := userRepository.db.Where("status = true").Order("updated_at desc").Find(&users).Error
	return users, err
}

func (userRepository *UserRepository) Create(user *models.User) error {
	err := userRepository.db.Create(user).Error
	return err
}
