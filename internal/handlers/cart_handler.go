// internal/handlers/cart_handler.go

package handlers

import (
	"net/http"
	"strconv"

	"github.com/DeviAnggita/my-online-store/internal/services"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService services.CartService
}

func NewCartHandler(cartService services.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

// POST /api/v1/cart/add
func (ch *CartHandler) AddToCart(c *gin.Context) {
	// Ambil data pelanggan dari konteks atau sesi (anda dapat menggunakan autentikasi JWT, cookie, dll.)
	customerID := uint(1) // Gantilah dengan cara yang sesuai untuk mendapatkan ID pelanggan

	// Ambil data produk dari permintaan
	var productData struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&productData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk menambahkan produk ke keranjang belanja
	err := ch.cartService.AddToCart(customerID, productData.ProductID, productData.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})
}

// GET /api/v1/cart
func (ch *CartHandler) GetCart(c *gin.Context) {
	// Retrieve customer ID from the authentication or session
	// For simplicity, let's assume the customer ID is provided in the URL
	customerID, err := strconv.ParseUint(c.Param("customerID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	cartItems, err := ch.cartService.GetCartItems(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

// // POST /api/v1/cart/add
// func (ch *CartHandler) AddToCart(c *gin.Context) {
// 	var cart models.Cart
// 	if err := c.BindJSON(&cart); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
// 		return
// 	}

// 	// Retrieve customer ID from the authentication or session
// 	// For simplicity, let's assume the customer ID is provided in the URL
// 	customerID, err := strconv.ParseUint(c.Param("customerID"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
// 		return
// 	}

// 	cart.CustomerID = uint(customerID)

// 	if err := ch.cartService.AddToCart(&cart); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "Product added to the cart"})
// }

// DELETE /api/v1/cart/:cartID
func (ch *CartHandler) DeleteFromCart(c *gin.Context) {
	cartID, err := strconv.ParseUint(c.Param("cartID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	if err := ch.cartService.DeleteFromCart(uint(cartID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted from the cart"})
}
