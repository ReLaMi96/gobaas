package handlers

import (
	"github.com/ReLaMi96/gobaas/sql"
	"github.com/ReLaMi96/gobaas/templates"
	"github.com/ReLaMi96/gobaas/utils"
	"github.com/ReLaMi96/gobaas/view"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BaseHandler struct {
	DB *gorm.DB
}

func (h BaseHandler) Base(c echo.Context) error {

	dbdetails, err := sql.GetDBdetails(h.DB)
	if err != nil {
		return err
	}

	queryStats, err := sql.QueryPerfRead(*h.DB)
	if err != nil {
		return err
	}

	return utils.Render(c, templates.Base(view.Dashboard(*dbdetails, queryStats)))
}
