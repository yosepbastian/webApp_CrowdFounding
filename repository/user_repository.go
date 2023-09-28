package repository

import (
	"web-app-crowdfounding/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Save(user models.User) (models.User, error)
	GetAllUser() ([]models.User, error)
	FindByEmail(email string) (models.User, error)
}

func (u *userRepository) GetAllUser() ([]models.User, error) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) Save(user models.User) (models.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := u.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
