package data

type WoolWorthsItemResponse struct {
	ItemId string `json:"sku"`
	Name   string `json:"name"`
	Price  Price  `json:"price"`
}
