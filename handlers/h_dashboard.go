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

	dbdetails, err := sql.GetDBdetails(h.DB)
	if err != nil {
		return err
	}

	queryStats, err := sql.QueryPerfRead(*h.DB)
	if err != nil {
		return err
	}

	return utils.Render(c, view.Dashboard(*dbdetails, queryStats))
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
