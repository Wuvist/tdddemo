package furycounter

import (
	"net/http"

	"github.com/labstack/echo"
)

// Bind apis to echo server
func Bind(e *echo.Echo) error {
	e.GET("/", hello)

	return nil
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func logout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
