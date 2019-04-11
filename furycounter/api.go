package furycounter

import (
	"net/http"

	"github.com/labstack/echo"
)

var counters map[string]*Counter

// Bind apis to echo server
func Bind(e *echo.Echo) error {
	e.GET("/", hello)

	counters = make(map[string]*Counter)

	e.POST("/api/furyCount.Add", add)
	e.POST("/api/furyCount.Hit", hit)
	e.POST("/api/furyCount.Block", block)
	e.POST("/api/furyCount.Get", get)

	return nil
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
