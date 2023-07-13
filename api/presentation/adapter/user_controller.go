package adapter

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/travel-buddy/domain"
	"github.com/Tomoya185-miyawaki/travel-buddy/usecase"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Login(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) Login(c echo.Context) error {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	_, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
