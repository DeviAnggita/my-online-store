package main

import (
	"github.com/DeviAnggita/my-online-store/config"
	"github.com/DeviAnggita/my-online-store/db/migrations"
	"github.com/DeviAnggita/my-online-store/internal/handlers"
	"github.com/DeviAnggita/my-online-store/internal/repository"
	"github.com/DeviAnggita/my-online-store/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Setup Database
	db, err := gorm.Open("mysql", config.AppConfig.DBUsername+":"+config.AppConfig.DBPassword+"@tcp("+config.AppConfig.DBHost+":"+config.AppConfig.DBPort+")/"+config.AppConfig.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Run Migrations
	migrations.RunMigrations(db)

	// // Initialize repository
	// productRepo := repository.NewProductRepository(db) // Make sure to replace db with your actual database connection

	// // Initialize service
	// productService := services.NewProductService(productRepo)

	// // Initialize handler
	// productHandler := handlers.NewProductHandler(productService)

	// Initialize repositories, services, and handlers
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	customerRepo := repository.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepo)
	customerHandler := handlers.NewCustomerHandler(customerService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	cartRepo := repository.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	// Initialize Gin router
	r := gin.Default()

	// Define API routes
	v1 := r.Group("/api/v1")
	{
		// Product Endpoints
		v1.GET("/products", productHandler.GetProducts)
		v1.GET("/products/:id", productHandler.GetProductByID)
		v1.POST("/products", productHandler.AddProduct)
		v1.PUT("/products/:id", productHandler.UpdateProduct)
		v1.DELETE("/products/:id", productHandler.DeleteProduct)

		// New route for getting products by category
		v1.GET("/products/category/:category", productHandler.GetProductsByCategory)

		// Customer Endpoints
		v1.POST("/register", customerHandler.RegisterCustomer)
		v1.POST("/login", customerHandler.LoginCustomer)

		// Order Endpoints
		v1.POST("/checkout", orderHandler.Checkout)

		// Cart Endpoints
		v1.GET("/cart", cartHandler.GetCart)
		v1.POST("/cart/add", cartHandler.AddToCart)
		v1.DELETE("/cart/:productID", cartHandler.DeleteFromCart)
	}

	// // Setup Redis
	// // redisClient := redis.NewClient(&redis.Options{
	// //     Addr:     config.AppConfig.RedisAddr,
	// //     Password: config.AppConfig.RedisPass,
	// //     DB:       0,
	// // })

	// Run the application on port 8080
	r.Run(":8080")
}
