package main

import (
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func handleEcho(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.Blob(200, c.Request().Header.Get("Content-Type"), body)
}
