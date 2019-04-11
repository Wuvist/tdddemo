package furycounter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func get(c echo.Context) error {
	name := c.Request().PostFormValue("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Must provide counter name")
	}

	counter, ok := counters[name]
	if !ok {
		return c.String(http.StatusBadRequest, "counter not found")
	}

	return c.String(http.StatusOK, strconv.Itoa(counter.Fury))
}
