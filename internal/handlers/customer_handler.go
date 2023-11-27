// internal/handlers/customer_handler.go

package handlers

import (
	"net/http"

	"github.com/DeviAnggita/my-online-store/internal/models"
	"github.com/DeviAnggita/my-online-store/internal/services"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService}
}

// POST /api/v1/register
func (ch *CustomerHandler) RegisterCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := ch.customerService.RegisterCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer registered successfully"})
}

// POST /api/v1/login
func (ch *CustomerHandler) LoginCustomer(c *gin.Context) {
	var loginInfo models.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	token, err := ch.customerService.LoginCustomer(loginInfo.Username, loginInfo.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
