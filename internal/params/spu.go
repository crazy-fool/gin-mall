package params

type EditSpuParam struct {
	Id         int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Img        string `json:"img" binding:"required"`
	Status     int    `json:"status" binding:"required"`
	BrandBn    string `json:"brand_bn" binding:"required"`
	StoreBn    string `json:"store_bn" binding:"required"`
	PriceRange string `json:"price_range" binding:"required"`
}
