package main

import (
	"example.com/mod/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/yamabiko", handler.YababikoAPI())
	// e.OPTIONS("/yababiko", handler.OptionsCheck())

	e.Logger.Fatal(e.Start(":8000"))
}