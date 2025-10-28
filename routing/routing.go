package routing

import (
	"github.com/ReLaMi96/gobaas/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetRoutes(app *echo.Echo, db *gorm.DB) {

	authHandler := handlers.UserAuthHandler{DB: db}
	baseHandler := handlers.BaseHandler{DB: db}
	dashHandler := handlers.DashboardHandler{DB: db}

	protected := app.Group("")
	protected.Use(authHandler.SessionHandler)

	app.GET("/login", authHandler.LoginForm)
	app.GET("/create-account-form", authHandler.CreateAccountForm)
	app.POST("/create-account", authHandler.CreateAccount)
	app.POST("/login-try", authHandler.Login)

	protected.GET("/", baseHandler.Base)
	protected.GET("/dashboard", dashHandler.Dashboard)
	protected.GET("/stat/all", dashHandler.Stats)
}
