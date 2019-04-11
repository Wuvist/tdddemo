package main

import (
	"github.com/Wuvist/tdddemo/furycounter"
	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()
	furycounter.Bind(e)

	// Start server
	e.Logger.Fatal(e.Start(":8323"))
}
