package sql

import (
	"github.com/ReLaMi96/gobaas/models"
	"gorm.io/gorm"
)

func SchemaStats(db gorm.DB) ([]models.SingleStat, error) {
	var result []models.SingleStat
	stat, err := TableCount(db)
	if err == nil {
		result = append(result, stat)
	}
	return result, err
}

func TableCount(db gorm.DB) (models.SingleStat, error) {
	var result models.SingleStat
	err := db.Raw("SELECT '' as title, COUNT(*) as value, '' as description FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';").Scan(&result).Error
	result.Title = "Total Tables"
	result.Description = "Number of tables in the database"

	return result, err
}
