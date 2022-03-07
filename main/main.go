package main

import (
	"tourism/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	routes.Register(app)

	app.Logger.Fatal(app.Start(":8080"))
}
