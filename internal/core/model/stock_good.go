package model

// StockGood 库存商品
type StockGood struct {
	Title            string           `json:"title,omitempty"`
	SkuCode          string           `json:"skuCode,omitempty"`
	Length           float64          `json:"length,omitempty"`
	Width            float64          `json:"width,omitempty"`
	Height           float64          `json:"height,omitempty"`
	Weight           float64          `json:"weight,omitempty"`
	ImgUrl           string           `json:"imgUrl,omitempty"`
	Cnt              int              `json:"cnt,omitempty"`
	StockGoodBatches []StockGoodBatch `json:"stockGoodBatches,omitempty"`
}

// StockGoodBatch 库存商品批次
type StockGoodBatch struct {
	BatchNumber    string `json:"batchNumber,omitempty"`
	ProductionDate string `json:"productionDate,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	Cnt            int    `json:"cnt,omitempty"`
}
