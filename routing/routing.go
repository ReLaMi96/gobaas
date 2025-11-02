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
	tableHandler := handlers.TableHandler{DB: db}

	protected := app.Group("")
	protected.Use(authHandler.SessionHandler)

	app.GET("/login", authHandler.LoginForm)
	app.GET("/create-account-form", authHandler.CreateAccountForm)
	app.POST("/create-account", authHandler.CreateAccount)
	app.POST("/login-try", authHandler.Login)

	protected.GET("/", baseHandler.BaseDashboard)
	protected.GET("/dashboard", dashHandler.Dashboard)
	protected.GET("/tables", tableHandler.Tables)

	protected.GET("/stat/all", dashHandler.Stats)
	protected.GET("/stat/top-queries", dashHandler.TopQueryList)
	protected.GET("/stat/schema-stats", dashHandler.SchemaStats)

	protected.GET("/tables/table-list", tableHandler.TableList)
	protected.GET("/tables/column-list", tableHandler.ColumnList)
}
