package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// show api --> return json
	// url --> '/users/:id"
	e.GET("users/:id", func(c echo.Context) error {
		jsonMap := map[string]string{
			"name": "example",
			"email": "example@example.com",
		}
		return c.JSON(http.StatusOK, jsonMap)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
