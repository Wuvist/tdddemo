package furycounter

import (
	"net/http"

	"github.com/labstack/echo"
)

func hit(c echo.Context) error {
	name := c.Request().PostFormValue("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Must provide counter name")
	}

	counter, ok := counters[name]
	if !ok {
		return c.String(http.StatusBadRequest, "counter not found")
	}

	counter.Hit()

	return c.String(http.StatusOK, "")
}
