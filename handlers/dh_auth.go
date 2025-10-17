package handlers

import (
	"errors"
	"time"

	"github.com/ReLaMi96/gobaas/models"
	"gorm.io/gorm"
)

func UserRead(d models.User, db gorm.DB) ([]models.User, error) {
	var result []models.User
	err := db.Select("id, email, username").Where("id = ? OR email = ?", d.ID, d.Email).Find(&result).Error
	return result, err
}

func UserReadRow(d models.User, db gorm.DB) (models.User, error) {
	var result models.User
	err := db.Select("id, email, username, password").Where("id = ? OR email = ?", d.ID, d.Email).First(&result).Error
	return result, err
}

func UserWrite(d models.User, db gorm.DB) *gorm.DB {
	return db.Create(&d)
}

func SessionReadRow(d models.Session, db gorm.DB) (models.Session, error) {
	var result models.Session
	err := db.Select("id, sessionkey, user_id").Where("sessionkey = ? and expiration > ? and last_use > ?", d.Sessionkey, time.Now(), time.Now().Add(-1*time.Hour)).First(&result).Error
	return result, err
}

func SessionWrite(d models.Session, db gorm.DB) error {
	result := db.Create(&d)
	if result == nil {
		return errors.New("cannot create session token")
	}
	return nil
}

func SessionDelete(d models.Session, db gorm.DB) (models.Session, error) {
	return models.Session{}, nil
}

func SessionUpdate(d models.Session, db gorm.DB) {
	db.Save(&d)
}
