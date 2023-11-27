package models

type Cart struct {
	ID         uint `gorm:"primary_key" json:"id"`
	CustomerID uint `json:"customer_id"`
	ProductID  uint `json:"product_id"`
	Quantity   int  `json:"quantity"`
}
