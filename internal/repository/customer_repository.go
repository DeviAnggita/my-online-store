// internal/repositories/customer_repository.go

package repository

import (
	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/jinzhu/gorm"
)

type CustomerRepository interface {
	RegisterCustomer(customer *models.Customer) error
	LoginCustomer(username, password string) (string, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *customerRepository {
	return &customerRepository{db}
}

func (cr *customerRepository) RegisterCustomer(customer *models.Customer) error {
	// Implement the registration logic using the database
	if err := cr.db.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func (cr *customerRepository) LoginCustomer(username, password string) (string, error) {
	// Implement the login logic using the database
	// Return a token or an error based on the authentication result
	// Note: This is a simplified example, you should use a proper authentication mechanism
	return "dummyToken", nil
}
