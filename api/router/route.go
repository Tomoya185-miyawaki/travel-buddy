package router

import (
	"fmt"
	"net/http"

	"github.com/Tomoya185-miyawaki/travel-buddy/helper"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	api := e.Group("/api/")
	api.GET("login", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	e.Start(fmt.Sprintf(":%s", helper.LoadEnv()))
}
