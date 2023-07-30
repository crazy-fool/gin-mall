package model

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	Name        string `gorm:"not null"`
	Code        string `gorm:"not null"`
	Account     string `gorm:"not null"`
	Mobile      string `gorm:"not null"`
	Password    string `gorm:"not null"`
	LastLoginAt time.Time
	gorm.Model
}

func (u *Customer) TableName() string {
	return "customers"
}
