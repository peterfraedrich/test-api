package main

import (
	"math/rand"
	"strconv"

	"github.com/labstack/echo/v4"
)

func handlePayloadSmall(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"size":    "1k",
		"payload": handlePayload(1),
	})
}

func handlePayloadMedium(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"size":    "750k",
		"payload": handlePayload(250),
	})
}

func handlePayloadLarge(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"size":    "1024",
		"payload": handlePayload(1024),
	})
}

func handlePayloadCustom(c echo.Context) error {
	size, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"size":    c.Param("size") + "k",
		"payload": handlePayload(size),
	})
}

func handlePayload(size int) []byte {
	payload := make([]byte, size*1000)
	rand.Read(payload)
	return payload
}
