// internal/services/order_service.go

package services

import (
	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/DeviAnggita/my-online-store/internal/repository"
)

type OrderService interface {
	// Checkout
	GetOrderByID(id uint) (*models.Order, error)
	AddOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{orderRepo}
}

func (os *orderService) GetOrderByID(id uint) (*models.Order, error) {
	return os.orderRepo.GetOrderByID(id)
}

func (os *orderService) AddOrder(order *models.Order) error {
	return os.orderRepo.AddOrder(order)
}

func (os *orderService) UpdateOrder(order *models.Order) error {
	return os.orderRepo.UpdateOrder(order)
}

func (os *orderService) DeleteOrder(id uint) error {
	return os.orderRepo.DeleteOrder(id)
}
