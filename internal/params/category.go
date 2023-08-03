package params

type CategoryEditParam struct {
	Id       uint   `json:"id" `
	Name     string `json:"name" binding:"required"`
	ParentId uint   `json:"parent_id"`
	IsParent uint8  `json:"is_parent"`
	Sort     uint   `json:"sort" `
}

type CategoryListParam struct {
	ParentId *uint `json:"parent_id" form:"parent_id" `
	PageSt
}
