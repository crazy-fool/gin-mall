package model

type Category struct {
	Name      string    `gorm:"not null" json:"name"`
	ParentId  uint      `gorm:"not null" json:"parent_id"`
	IsParent  uint8     `gorm:"not null" json:"is_parent"`
	Sort      uint      `gorm:"not null" json:"sort"`
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt LocalTime `json:"created_at"`
	UpdatedAt LocalTime `json:"updated_at"`
}

func (u *Category) TableName() string {
	return "hmh_category"
}
