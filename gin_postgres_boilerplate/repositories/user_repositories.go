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

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	err := ur.db.
		Select("id", "name", "email", "status", "created_at", "updated_at").
		Where("status = true").
		Order("updated_at desc").
		Find(&users).
		Error
	return users, err
}

func (ur *UserRepository) Create(user *models.User) error {
	err := ur.db.Create(user).Error
	return err
}

func (ar *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ar.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
