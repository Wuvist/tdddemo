package main

import (
	"github.com/Wuvist/tdddemo/newtrans"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()
	newtrans.Bind(e)

	// Start server
	e.Logger.Fatal(e.Start(":8323"))
}
