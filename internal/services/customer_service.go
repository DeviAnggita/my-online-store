// internal/services/customer_service.go

package services

import (
	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/DeviAnggita/my-online-store/internal/repository"
)

type CustomerService interface {
	RegisterCustomer(customer *models.Customer) error
	LoginCustomer(username, password string) (string, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo}
}

func (cs *customerService) RegisterCustomer(customer *models.Customer) error {
	// Implement the registration logic using the repository
	return cs.repo.RegisterCustomer(customer)
}

func (cs *customerService) LoginCustomer(username, password string) (string, error) {
	// Implement the login logic using the repository
	return cs.repo.LoginCustomer(username, password)
}
