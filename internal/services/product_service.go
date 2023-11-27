// internal/services/product_service.go

package services

import (
	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/DeviAnggita/my-online-store/internal/repository"
)

type ProductService interface {
	GetProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	AddProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	GetProductsByCategory(category string) ([]models.Product, error)
}

type productService struct {
	repo repository.ProductRepository // Corrected type name
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

// Implement other methods...

func (ps *productService) GetProductsByCategory(category string) ([]models.Product, error) {
	// Implement logic to get products by category from the repository
	// For example:
	// return s.repo.GetProductsByCategory(category)
	return nil, nil
}

func (ps *productService) GetProducts() ([]models.Product, error) {
	return ps.repo.GetAllProducts()
}

func (ps *productService) GetProductByID(id uint) (*models.Product, error) {
	return ps.repo.GetProductByID(id)
}

func (ps *productService) AddProduct(product *models.Product) error {
	return ps.repo.AddProduct(product)
}

func (ps *productService) UpdateProduct(product *models.Product) error {
	return ps.repo.UpdateProduct(product)
}

func (ps *productService) DeleteProduct(id uint) error {
	return ps.repo.DeleteProduct(id)
}
