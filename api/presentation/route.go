package presentation

import (
	"fmt"

	"github.com/Tomoya185-miyawaki/travel-buddy/helper"
	"github.com/Tomoya185-miyawaki/travel-buddy/presentation/adapter"
	"github.com/Tomoya185-miyawaki/travel-buddy/presentation/middleware"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func Route(uc adapter.IUserController) {
	e := echo.New()
	e.Use(middleware.Log())
	e.Use(m.Recover())

	api := e.Group("/api/")
	api.POST("login", uc.Login)

	e.Start(fmt.Sprintf(":%s", helper.LoadEnv()))
}
