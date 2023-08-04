package params

type CategoryEditParam struct {
	Id       int32  `json:"id" `
	Name     string `json:"name" binding:"required"`
	ParentId *int32 `json:"parent_id"`
	IsParent *int8  `json:"is_parent"`
	Sort     *int32 `json:"sort" `
}

type CategoryListParam struct {
	ParentId *int32 `json:"parent_id" form:"parent_id" `
	*PageSt
}
