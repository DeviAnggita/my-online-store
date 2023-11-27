package models

type Product struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}
