package params

type SpecEditParam struct {
	Id         int32  `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	CategoryId int32  `json:"category_id" binding:"required"`
	GroupId    int32  `json:"group_id" binding:"required"`
	IsGenerate *int8  `json:"is_generate" binding:"required"`
	Searching  *int8  `json:"searching" binding:"required"`
}

type SpecListParam struct {
	PageSt
}

type GroupEditParam struct {
	Id         int32  `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	CategoryId int32  `json:"category_id" binding:"required"`
}

type GroupListParam struct {
	PageSt
}
