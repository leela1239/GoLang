package models

type Product struct {
	ProductId   string    `json:"id"`
	ProductName string    `json:"name"`
	Variants    []Variant `json:"variants"`
}

type Variant struct {
	Variantid string  `json:"id"`
	ProductId string  `json:"pid"`
	Size      string  `json:"size"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
}
