package model

type Tick struct {
	Instrument string `json:"instrument"`
	Price1000  int    `json:"price_1000"`
}
