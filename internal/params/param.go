package params

type PageSt struct {
	Page     int `json:"page" uri:"page" form:"page"`
	PageSize int `json:"page_size"  uri:"page_size" form:"page_size"`
}
