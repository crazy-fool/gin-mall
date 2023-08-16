// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSkuDetail = "hmh_sku_detail"

// SkuDetail mapped from table <hmh_sku_detail>
type SkuDetail struct {
	ID        int32      `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	SkuID     *int32     `gorm:"column:sku_id;type:int(11)" json:"sku_id"`
	Images    *string    `gorm:"column:images;type:varchar(2000)" json:"images"`
	Indexes   *string    `gorm:"column:indexes;type:varchar(255)" json:"indexes"`
	OwnSpec   *string    `gorm:"column:own_spec;type:varchar(2000)" json:"own_spec"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

// TableName SkuDetail's table name
func (*SkuDetail) TableName() string {
	return TableNameSkuDetail
}
