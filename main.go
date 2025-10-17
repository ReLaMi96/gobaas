package main

import (
	"github.com/ReLaMi96/gobaas/routing"
	"github.com/ReLaMi96/gobaas/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	dsn := "host=localhost user=gobaas password=gobaas dbname=gobaas port=5432 sslmode=disable TimeZone=Europe/Budapest"

	db, err := utils.DBinit(dsn)
	if err != nil {
		panic(err)
	}

	if err := utils.AutoMigrate(db); err != nil {
		panic(err)
	}

	println("> Base tables migrated")

	app := echo.New()

	app.Static("/static", "static")

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	routing.SetRoutes(app, db)

	app.Start(":9696")

}
