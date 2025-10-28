package utils

import (
	"fmt"

	"github.com/ReLaMi96/gobaas/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBdetails struct {
	Status    string
	Host      string
	Port      string
	DBname    string
	DBversion string
	SSLmode   string
	Uptime    string
	CPU       string
	RAM       string
	Space     string
}

func DBinit(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqldb.SetMaxOpenConns(99)
	sqldb.SetMaxIdleConns(1)

	fmt.Println("> Database connected")

	EnableExtensions(db, "pg_stat_statements")

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Session{},
	)
}

func EnableExtensions(db *gorm.DB, ext string) error {
	sql := fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS %s;", ext)
	err := db.Exec(sql).Error
	return err
}
