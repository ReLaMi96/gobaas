package utils

import (
	"fmt"

	"github.com/ReLaMi96/gobaas/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Session{},
	)
}
