package main

import (
	"fmt"
	"net/http"

	"product-res-api/config"
	"product-res-api/pkg/product"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Default route handler
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})

	// Connect to the database
	db := config.InitDB()
	defer config.CloseDB(db)

	db.AutoMigrate(&product.Product{})
	fmt.Println("AutoMigrate Product...")
	fmt.Println("Database connected")

	// Group routes related to product endpoints
	setupProductRoutes(e)

	// Start the server
	apiPort := config.ApiPort()
	e.Logger.Fatal(e.Start(":" + apiPort))
}

// setupProductRoutes เป็นฟังก์ชันที่ใช้สำหรับกำหนดเส้นทางที่เกี่ยวข้องกับ endpoints ของสินค้า
func setupProductRoutes(e *echo.Echo) {
	// Group routes related to product endpoints
	productRoute := e.Group("/product")
	productRoute.GET("", product.GetProducts)    // Get all products
	productRoute.GET("/:id", product.GetProduct) // Get a specific product by ID

	productRoute.POST("", product.CreateProduct)       // Create a new product
	productRoute.DELETE("/:id", product.DeleteProduct) // Delete a product by ID
	productRoute.PUT("/:id", product.UpdateProduct)    // Update a product by ID
}
