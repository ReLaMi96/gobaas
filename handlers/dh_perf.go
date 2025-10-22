package handlers

import (
	"github.com/ReLaMi96/gobaas/components"
	"gorm.io/gorm"
)

func QueryPerfRead(db gorm.DB) ([]components.QueryPerf, error) {
	var result []components.QueryPerf
	err := db.Raw("SELECT query, calls, total_exec_time, mean_exec_time, rows FROM pg_stat_statements ORDER BY queryid DESC LIMIT 100;").Scan(&result).Error
	return result, err
}
