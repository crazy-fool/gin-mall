package params

type EditCategoryParam struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	ParentId int    `json:"parent_id" binding:"required"`
	IsParent int    `json:"is_parent" binding:"required"`
	Sort     int    `json:"sort" binding:"required"`
}
