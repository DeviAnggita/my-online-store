// internal/handlers/product_handler.go

package handlers

import (
	"net/http"
	"strconv"

	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/DeviAnggita/my-online-store/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

// Implement handler methods...

// GET /api/v1/products/category/:category
func (ph *ProductHandler) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	products, err := ph.productService.GetProductsByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) GetProducts(c *gin.Context) {
	products, err := ph.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := ph.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := ph.productService.AddProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	updatedProduct.ID = uint(id)
	if err := ph.productService.UpdateProduct(&updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := ph.productService.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
