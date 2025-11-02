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

	stat, err = DBSize(db)
	if err == nil {
		result = append(result, stat)
	}

	stat, err = ActiveConnections(db)
	if err == nil {
		result = append(result, stat)
	}

	return result, err
}

func TableCount(db gorm.DB) (models.SingleStat, error) {
	var result models.SingleStat
	err := db.Raw("SELECT '' as title, COUNT(*) as value, '' as description FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';").Scan(&result).Error
	result.Title = "Total Tables"
	result.Description = "Number of tables"

	return result, err
}

func DBSize(db gorm.DB) (models.SingleStat, error) {
	var result models.SingleStat
	err := db.Raw("SELECT '' as title, pg_size_pretty(pg_database_size(current_database())) as value, '' as description").Scan(&result).Error
	result.Title = "Database Size"
	result.Description = "Total size of the database"

	return result, err
}

func ActiveConnections(db gorm.DB) (models.SingleStat, error) {
	var result models.SingleStat
	err := db.Raw("SELECT '' as title, COUNT(*)::text as value, '' as description FROM pg_stat_activity WHERE datname = current_database()").Scan(&result).Error
	result.Title = "Active Connections"
	result.Description = "Current database connections"

	return result, err
}

func TableList(db gorm.DB, search string) (models.TableData, error) {
	columns := []string{"Schema", "Name", "Type"}
	rows := []models.TableRow{}
	var schema, name, tableType string

	search = "%" + search + "%"

	sqlrows, err := db.Raw(`
		SELECT 
			table_schema,
			table_name,
			table_type
		FROM information_schema.tables 
		WHERE table_schema = 'public'
		AND table_type = 'BASE TABLE'
		AND table_name LIKE ?
		ORDER BY table_name;
	`, search).Rows()

	if err != nil {
		return models.TableData{}, err
	}

	for sqlrows.Next() {
		err = sqlrows.Scan(&schema, &name, &tableType)
		if err != nil {
			return models.TableData{}, err
		}
		row := models.TableRow{
			Cells: []string{schema, name, tableType},
		}
		rows = append(rows, row)
	}

	result := models.TableData{
		Columns: columns,
		Rows:    rows,
	}
	return result, nil
}

func ColumnList(db gorm.DB, table string, schema string) (models.TableData, error) {
	columns := []string{"Name", "Type", "Nullable"}
	rows := []models.TableRow{}
	var col1, col2, col3 string

	sqlrows, err := db.Raw(`
		SELECT 
 		    column_name,
   	 		data_type,
    		is_nullable
		FROM information_schema.columns 
		WHERE table_schema = ? 
		AND table_name = ?
		ORDER BY ordinal_position;
	`, schema, table).Rows()

	if err != nil {
		return models.TableData{}, err
	}

	for sqlrows.Next() {
		err = sqlrows.Scan(&col1, &col2, &col3)
		if err != nil {
			return models.TableData{}, err
		}
		row := models.TableRow{
			Cells: []string{col1, col2, col3},
		}
		rows = append(rows, row)
	}

	result := models.TableData{
		Columns: columns,
		Rows:    rows,
	}
	return result, nil
}
