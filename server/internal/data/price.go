package data

type Price struct {
	OriginalPrice  float32 `json:"originalPrice"`
	SalePrice      float32 `json:"salePrice"`
	SavePrice      float32 `json:"savePrice"`
	SavePercentage float32 `json:"savePercentage"`
}
