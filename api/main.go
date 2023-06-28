package main

import (
	"github.com/Tomoya185-miyawaki/travel-buddy/middleware"
	"github.com/Tomoya185-miyawaki/travel-buddy/router"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Log())
	e.Use(m.Recover())
	router.Route(e)
}
