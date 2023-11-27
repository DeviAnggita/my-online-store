package migrations

import (
	"github.com/DeviAnggita/my-online-store/internal/models"

	"github.com/jinzhu/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Cart{})
}
