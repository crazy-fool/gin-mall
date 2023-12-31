// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSpu = "hmh_spu"

// Spu mapped from table <hmh_spu>
type Spu struct {
	ID         int32      `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	Spu        string     `gorm:"column:spu;type:varchar(20);not null;comment:spucode" json:"spu"`          // spucode
	Name       string     `gorm:"column:name;type:varchar(50);not null;comment:商品名称" json:"name"`           // 商品名称
	Img        *string    `gorm:"column:img;type:varchar(255);comment:图片" json:"img"`                       // 图片
	Status     *bool      `gorm:"column:status;type:tinyint(1);default:1;comment:状态1启用，0、禁用" json:"status"` // 状态1启用，0、禁用
	BrandBn    *string    `gorm:"column:brand_bn;type:varchar(20);comment:品牌code" json:"brand_bn"`          // 品牌code
	StoreBn    *string    `gorm:"column:store_bn;type:varchar(20);comment:商户号" json:"store_bn"`             // 商户号
	SaleCnt    *int32     `gorm:"column:sale_cnt;type:int(11);comment:销量" json:"sale_cnt"`                  // 销量
	PriceRange *string    `gorm:"column:price_range;type:varchar(255);comment:价格区间" json:"price_range"`     // 价格区间
	CateID     *int32     `gorm:"column:cate_id;type:int(11);comment:子分类id" json:"cate_id"`                 // 子分类id
	OneCateID  *int32     `gorm:"column:one_cate_id;type:int(11);comment:一级分类id" json:"one_cate_id"`        // 一级分类id
	TwoCateID  *int32     `gorm:"column:two_cate_id;type:int(11);comment:二级分类id" json:"two_cate_id"`        // 二级分类id
	CreatedAt  *time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`           // 创建时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`           // 更新时间
}

// TableName Spu's table name
func (*Spu) TableName() string {
	return TableNameSpu
}
