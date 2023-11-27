// internal/repository/order_repository.go

package repository

import (
	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	// checkout
	GetOrderByID(id uint) (*models.Order, error)
	AddOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (or *orderRepository) GetOrderByID(id uint) (*models.Order, error) {
	// Implement logic to fetch order by ID from the database
	return nil, nil
}

func (or *orderRepository) AddOrder(order *models.Order) error {
	// Implement logic to add order to the database
	return nil
}

func (or *orderRepository) UpdateOrder(order *models.Order) error {
	// Implement logic to update order in the database
	return nil
}

func (or *orderRepository) DeleteOrder(id uint) error {
	// Implement logic to delete order from the database
	return nil
}
