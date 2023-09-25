package usecase

import (
	"web-app-crowdfounding/models"
	"web-app-crowdfounding/repository"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

type UserUseCase interface {
	GetAllUser() ([]models.User, error)
	RegisterUser(input models.RegisterUserInput) (models.User, error)
}

func (u *userUseCase) GetAllUser() ([]models.User, error) {
	return u.userRepo.GetAllUser()
}

func (u *userUseCase) RegisterUser(input models.RegisterUserInput) (models.User, error) {
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.Role = "user"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)

	newUser, err := u.userRepo.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func NewUserUseCase(uRepo repository.UserRepository) *userUseCase {
	return &userUseCase{uRepo}
}
