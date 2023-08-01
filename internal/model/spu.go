package model

import (
	"gorm.io/gorm"
)

type Spu struct {
	Name       string `gorm:"not null"`
	Spu        string `gorm:"not null"`
	Img        string `gorm:"not null"`
	Status     int    `gorm:"not null"`
	BrandBn    string `gorm:"not null"`
	StoreBn    string `gorm:"not null"`
	SaleCnt    int    `gorm:"default:0"`
	PriceRange string `gorm:"default:"`
	gorm.Model
}

func (u *Spu) TableName() string {
	return "hmh_spu"
}
