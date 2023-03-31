package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)


type YamabikoParam struct {
	Message string `json:"message"`
}

func YababikoAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(YamabikoParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"Hello": param.Message})
	}
}