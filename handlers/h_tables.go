package handlers

import (
	"github.com/ReLaMi96/gobaas/components"
	"github.com/ReLaMi96/gobaas/sql"
	"github.com/ReLaMi96/gobaas/utils"
	"github.com/ReLaMi96/gobaas/view"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TableHandler struct {
	DB *gorm.DB
}

func (h TableHandler) Tables(c echo.Context) error {

	if c.Request().Header.Get("HX-Request") != "" {

		tables, err := sql.TableList(*h.DB, "")
		if err != nil {
			return err
		}

		columns, err := sql.ColumnList(*h.DB, "", "")
		if err != nil {
			return err
		}

		return utils.Render(c, view.Tables(tables, columns))
	}

	return BaseHandler{DB: h.DB}.BaseTables(c)
}

func (h TableHandler) TableList(c echo.Context) error {

	search := c.FormValue("table-search")

	data, err := sql.TableList(*h.DB, search)
	if err != nil {
		return err
	}

	return utils.Render(c, components.List(data))
}

func (h TableHandler) ColumnList(c echo.Context) error {

	tableName := c.FormValue("tableName")
	schema := c.FormValue("schema")

	data, err := sql.ColumnList(*h.DB, tableName, schema)
	if err != nil {
		return err
	}

	return utils.Render(c, components.ColumnList(data))
}
