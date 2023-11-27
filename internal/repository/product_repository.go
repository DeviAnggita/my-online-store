// internal/repository/product_repository.go

package repository

import (
	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	AddProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	GetProductsByCategory(category string) ([]models.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}

// Implement other methods...

func (pr *ProductRepositoryImpl) GetProductsByCategory(category string) ([]models.Product, error) {
	var products []models.Product
	if err := pr.db.Where("category = ?", category).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *ProductRepositoryImpl) GetAllProducts() ([]models.Product, error) {
	// Implement logic to fetch all products from the database
	return []models.Product{}, nil
}

func (pr *ProductRepositoryImpl) GetProductByID(id uint) (*models.Product, error) {
	// Implement logic to fetch a product by ID from the database
	return &models.Product{}, nil
}

func (pr *ProductRepositoryImpl) AddProduct(product *models.Product) error {
	// Implement logic to add a product to the database
	return nil
}

func (pr *ProductRepositoryImpl) UpdateProduct(product *models.Product) error {
	// Implement logic to update a product in the database
	return nil
}

func (pr *ProductRepositoryImpl) DeleteProduct(id uint) error {
	// Implement logic to delete a product from the database
	return nil
}
