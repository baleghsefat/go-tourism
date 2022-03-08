package routes

import (
	CompanyController "tourism/app/Http/Controllers/Company"
	UserController "tourism/app/Http/Controllers/User"

	"github.com/labstack/echo/v4"
)

func Register(app *echo.Echo) {
	userRoutes(app)
	companyRoutes(app)
}

func userRoutes(app *echo.Echo) {
	userGroup := app.Group("users")
	userGroup.GET("", UserController.Index)
	userGroup.GET("/:id", UserController.Show)
	userGroup.POST("", UserController.Store)
	userGroup.PUT("/:id", UserController.Update)
	userGroup.DELETE("/:id", UserController.Destroy)
}

func companyRoutes(app *echo.Echo) {
	companyGroup := app.Group("companies")
	companyGroup.GET("", CompanyController.Index)
	companyGroup.GET("/:id", CompanyController.Show)
	companyGroup.POST("", CompanyController.Store)
	companyGroup.PUT("/:id", CompanyController.Update)
	companyGroup.DELETE("/:id", CompanyController.Destroy)
}
