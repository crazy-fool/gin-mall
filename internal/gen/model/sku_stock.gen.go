// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSkuStock = "hmh_sku_stock"

// SkuStock mapped from table <hmh_sku_stock>
type SkuStock struct {
	ID        int32      `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	Sku       *string    `gorm:"column:sku;type:varchar(20)" json:"sku"`
	Stock     *int32     `gorm:"column:stock;type:int(11)" json:"stock"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

// TableName SkuStock's table name
func (*SkuStock) TableName() string {
	return TableNameSkuStock
}