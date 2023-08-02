// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSku = "hmh_sku"

// Sku mapped from table <hmh_sku>
type Sku struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string    `gorm:"column:name;not null;comment:skuname" json:"name"`                          // skuname
	Sku         string    `gorm:"column:sku;not null;comment:sku" json:"sku"`                                // sku
	Price       float64   `gorm:"column:price;not null;default:0.00;comment:售价" json:"price"`                // 售价
	MarketPrice float64   `gorm:"column:market_price;not null;default:0.00;comment:市场价" json:"market_price"` // 市场价
	Img         string    `gorm:"column:img;comment:图片" json:"img"`                                          // 图片
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Sku's table name
func (*Sku) TableName() string {
	return TableNameSku
}