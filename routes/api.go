package routes

import (
	UserController "tourism/app/Http/Controllers/User"

	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo) {
	app.GET("users", UserController.Index)
	app.GET("users/:id", UserController.Show)
	app.POST("users", UserController.Store)
	app.PUT("users/:id", UserController.Update)
	app.DELETE("users/:id", UserController.Destroy)
}
