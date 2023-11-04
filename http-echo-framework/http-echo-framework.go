package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// create a new echo
	ec := echo.New()

	// create middleware
	ec.Use(middleware.Logger())
	ec.Use(middleware.Recover())

	ec.GET("/", handleIndex)

	ec.Logger.Fatal(ec.Start(":1327"))
}

func handleIndex(ctxt echo.Context) error {
	return ctxt.String(http.StatusOK, "<h1>HELLO, MUM!</h1>")
}
