package main

import (
	"fmt"
	"net/http"

	"github.com/Tomoya185-miyawaki/travel-buddy/helper"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", helper.LoadEnv())))
}
