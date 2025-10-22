package handlers

import (
	"github.com/ReLaMi96/gobaas/components"
	"github.com/ReLaMi96/gobaas/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PerformanceHandler struct {
	DB *gorm.DB
}

func (h PerformanceHandler) QueryStats(c echo.Context) error {

	stats, err := QueryPerfRead(*h.DB)
	if err != nil {
		return err
	}

	return utils.Render(c, components.QueryStats(stats))
}
