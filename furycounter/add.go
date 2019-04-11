package furycounter

import (
	"net/http"

	"github.com/labstack/echo"
)

func add(c echo.Context) error {
	name := c.Request().PostFormValue("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Must provide counter name")
	}

	if _, ok := counters[name]; ok {
		return c.String(http.StatusBadRequest, "counter exist")
	}

	counter := &Counter{}
	counters[name] = counter

	return c.String(http.StatusOK, "")
}
