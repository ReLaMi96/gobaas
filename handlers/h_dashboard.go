package handlers

import (
	"github.com/ReLaMi96/gobaas/components"
	"github.com/ReLaMi96/gobaas/sql"
	"github.com/ReLaMi96/gobaas/utils"
	"github.com/ReLaMi96/gobaas/view"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DashboardHandler struct {
	DB *gorm.DB
}

func (h DashboardHandler) Dashboard(c echo.Context) error {

	if c.Request().Header.Get("HX-Request") != "" {

		queryStats, err := sql.QueryPerfRead(*h.DB)
		if err != nil {
			return err
		}

		return utils.Render(c, view.Dashboard(utils.DBdetails{}, queryStats, nil))
	}

	baseHandler := BaseHandler{DB: h.DB}
	return baseHandler.BaseDashboard(c)
}

func (h DashboardHandler) Stats(c echo.Context) error {
	result, err := sql.GetDBdetails(h.DB)
	if err != nil {
		return err
	}

	return utils.Render(c, components.DBdetails(*result))
}

func (h DashboardHandler) Status(c echo.Context) error {

	status, err := sql.CheckDatabaseHealth(h.DB)
	if err != nil {
		return err
	}

	return utils.Render(c, components.Status(status))
}

func (h DashboardHandler) TopQueryList(c echo.Context) error {

	queryStats, err := sql.QueryPerfRead(*h.DB)
	if err != nil {
		return err
	}

	return utils.Render(c, components.QueryStats(queryStats))
}

func (h DashboardHandler) SchemaStats(c echo.Context) error {

	statBoard, err := sql.SchemaStats(*h.DB)
	if err != nil {
		return err
	}

	return utils.Render(c, components.StatBoard(statBoard))
}
