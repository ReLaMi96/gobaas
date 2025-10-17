package handlers

import (
	"github.com/ReLaMi96/gobaas/templates"
	"github.com/ReLaMi96/gobaas/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DashboardHandler struct {
	DB *gorm.DB
}

func (h DashboardHandler) Dashboard(c echo.Context) error {
	return utils.Render(c, templates.Dashboard())
}
