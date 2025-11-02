package models

type TableData struct {
	Columns []string
	Rows    []TableRow
}

type TableRow struct {
	Cells []string
}
