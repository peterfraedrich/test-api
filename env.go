package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

func handleEnv(c echo.Context) error {
	env := os.Environ()
	return c.JSON(200, map[string]interface{}{
		"ENV": env,
	})
}
