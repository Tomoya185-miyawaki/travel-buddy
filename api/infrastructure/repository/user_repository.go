package repository

import (
	"github.com/Tomoya185-miyawaki/travel-buddy/domain"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmain(user *domain.User, email string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmain(user *domain.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}
