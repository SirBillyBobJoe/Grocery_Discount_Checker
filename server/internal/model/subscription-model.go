package model

type SubcriptionModel struct {
	ItemId        string   `bson:"_id" json:"itemId"`
	Emails        []string `bson:"emails" json:"emails"`
	OriginalPrice float32  `bson:"orignalPrice" json:"originalPrice"`
	CurrentPrice  float32  `bson:"currentPrice" json:"currentPrice"`
}
