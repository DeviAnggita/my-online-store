// internal/handlers/order_handler.go

package handlers

import (
	"net/http"
	"strconv"

	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/DeviAnggita/my-online-store/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{orderService}
}

// POST /api/v1/checkout
func (oh *OrderHandler) Checkout(c *gin.Context) {
	// Implementation for the checkout endpoint
	// You can use c.ShouldBindJSON to parse JSON payload if needed

	c.JSON(http.StatusCreated, gin.H{"message": "Checkout completed successfully"})
}

// GET /api/v1/order/:id
func (oh *OrderHandler) GetOrderByID(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := oh.orderService.GetOrderByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// POST /api/v1/order
func (oh *OrderHandler) AddOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := oh.orderService.AddOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order added successfully"})
}

// PUT /api/v1/order/:id
func (oh *OrderHandler) UpdateOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	order.ID = uint(orderID)

	if err := oh.orderService.UpdateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// DELETE /api/v1/order/:id
func (oh *OrderHandler) DeleteOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := oh.orderService.DeleteOrder(uint(orderID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
