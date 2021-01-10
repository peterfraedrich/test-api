package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// PING
	e.Any("/api/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"METOD": c.Request().Method,
		})
	})

	// UTILITY
	e.GET("/api/env", handleEnv)
	e.POST("/api/echo", handleEcho)

	// PAYLOAD
	e.GET("/api/payload/small", handlePayloadSmall)
	e.GET("/api/payload/medium", handlePayloadMedium)
	e.GET("/api/payload/large", handlePayloadLarge)
	e.GET("/api/payload/:size", handlePayloadCustom)

	// ROUTES
	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"routes": e.Routes(),
		})
	})

	// ACCESS LOGGING
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// SERVE
	e.Logger.Fatal(e.Start(":8080"))

}
