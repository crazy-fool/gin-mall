package params

type CategoryEditParam struct {
	Id       *uint  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	ParentId *uint  `json:"parent_id" binding:"required"`
	IsParent *uint8 `json:"is_parent" binding:"required"`
	Sort     *uint  `json:"sort" binding:"required"`
}

type CategoryListParam struct {
	PageSt
}
