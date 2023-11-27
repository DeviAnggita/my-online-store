package models

type Order struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	CustomerID  uint    `json:"customer_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
	OrderDate   string  `json:"order_date"`
}

// func (Order) TableName() string {
// 	return "orderr"
// }
