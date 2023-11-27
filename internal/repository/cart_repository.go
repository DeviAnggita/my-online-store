// internal/repositories/cart_repository.go

package repository

import (
	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/jinzhu/gorm"
)

type CartRepository interface {
	GetCartItems(customerID uint) ([]models.Cart, error)
	// GetCartByCustomerID(customerID uint) ([]models.Cart, error)
	AddCart(cart *models.Cart) error
	DeleteFromCart(cartID uint) error
}

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepo {
	return &cartRepo{db}
}

// GetCartItems retrieves the list of products in the shopping cart from the database.
func (cr *cartRepo) GetCartItems(customerID uint) ([]models.Cart, error) {
	var cartItems []models.Cart
	if err := cr.db.Where("customer_id = ?", customerID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}

// AddToCart adds a product to the shopping cart in the database.
func (cr *cartRepo) AddCart(cart *models.Cart) error {
	if err := cr.db.Create(cart).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFromCart deletes a product from the shopping cart in the database based on the product ID.
func (cr *cartRepo) DeleteFromCart(cartID uint) error {
	if err := cr.db.Delete(&models.Cart{}, cartID).Error; err != nil {
		return err
	}
	return nil
}

// // UpdateCart memperbarui jumlah produk di keranjang belanja
// func (r *cartRepo) UpdateCart(cart *models.Cart) error {
// 	// Implementasi untuk memperbarui jumlah produk di keranjang belanja di database
// 	// ...
// 	return nil
// }
