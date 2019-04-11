package furycounter

import (
	"net/http"

	"github.com/labstack/echo"
)

func block(c echo.Context) error {
	name := c.Request().PostFormValue("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Must provide counter name")
	}

	counter, ok := counters[name]
	if !ok {
		return c.String(http.StatusBadRequest, "counter not found")
	}

	counter.Block()

	return c.String(http.StatusOK, "")
}
