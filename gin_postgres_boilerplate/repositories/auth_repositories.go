package repositories

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (authRepository *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := authRepository.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
