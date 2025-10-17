package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint
	Email       string
	Date        time.Time
	Type        string
	Note        string
	Username    string
	Password    string
	Role        string
	Login       bool
	PasswordExp time.Time
}

type Session struct {
	gorm.Model
	ID         uint
	UserID     uint
	Sessionkey string
	Expiration time.Time
	LastUse    time.Time
	Refreshes  int
	ValidTo    time.Time
}
