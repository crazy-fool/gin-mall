package params

type PageSt struct {
	Page     int `json:"page" uri:"page" form:"page"`
	PageSize int `json:"page_size"  uri:"page_size" form:"page_size"`
}

func (p *PageSt) OffsetLimit() (int, int) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = 10
	}
	return (p.Page - 1) * p.PageSize, p.PageSize

}
