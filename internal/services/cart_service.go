// internal/services/cart_service.go

package services

import (
	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/DeviAnggita/my-online-store/internal/repository"
)

type CartService interface {
	GetCartItems(customerID uint) ([]models.Cart, error)
	AddToCart(customerID, productID uint, quantity int) error
	// AddToCart(cart *models.Cart) error
	DeleteFromCart(cartID uint) error
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo: repo}
}

// GetCart mendapatkan daftar produk yang telah ditambahkan ke keranjang belanja
func (cs *cartService) GetCartItems(customerID uint) ([]models.Cart, error) {
	// Implementasi untuk mendapatkan daftar produk dari repository
	return cs.repo.GetCartItems(customerID)
}

// AddToCart menambahkan produk ke keranjang belanja pelanggan
func (cs *cartService) AddToCart(customerID, productID uint, quantity int) error {
	// Anda dapat menambahkan logika validasi atau logika bisnis lainnya di sini
	return cs.repo.AddCart(&models.Cart{
		CustomerID: customerID,
		ProductID:  productID,
		Quantity:   quantity,
	})
}

// DeleteFromCart menghapus produk dari keranjang belanja
func (cs *cartService) DeleteFromCart(cartID uint) error {
	// Implementasi untuk menghapus produk dari keranjang belanja di repository
	return cs.repo.DeleteFromCart(cartID)
}

// // UpdateCart memperbarui jumlah produk di keranjang belanja
// func (s *cartService) UpdateCart(cartID, quantity int) error {
// 	// Implementasi untuk memperbarui jumlah produk di keranjang belanja di repository
// 	return nil
// }
