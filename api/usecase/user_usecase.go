package usecase

import (
	"github.com/Tomoya185-miyawaki/travel-buddy/domain"
	"github.com/Tomoya185-miyawaki/travel-buddy/infrastructure/repository"
	"github.com/Tomoya185-miyawaki/travel-buddy/presentation/validation"
)

type IUserUsecase interface {
	Login(user domain.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validation.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validation.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) Login(user domain.User) (string, error) {
	if err := uu.uv.UserValidator(user); err != nil {
		return "", err
	}
	storedUser := domain.User{}
	if err := uu.ur.GetUserByEmain(&storedUser, user.Email); err != nil {
		return "", err
	}
	return "", nil
}
