package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nuchit2019/product-res-api/config"
	"github.com/nuchit2019/product-res-api/product"
)

func main() {

	// Initialize Echo instance
	e := echo.New()

	// Default route handler
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})

	// Connect to the database
	config.InitDB()
	gorm := config.DB()

	// Retrieve the underlying *sql.DB instance from GORM
	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	// Ping the database to check connectivity
	dbGorm.Ping()

	// Group routes related to product endpoints
	productRoute := e.Group("/product")
	productRoute.GET("", product.GetProducts)    // Get all products
	productRoute.GET("/:id", product.GetProduct) // Get a specific product by ID

	productRoute.POST("", product.CreateProduct)    // Create a new product
	productRoute.DELETE("/:id", product.DeleteProduct) // Delete a product by ID
	productRoute.PUT("/:id", product.UpdateProduct)    // Update a product by ID

	// Start the server on port 8080
	apiPort:=config.ApiPort();
	// e.Logger.Fatal(e.Start(":8080"))
	e.Logger.Fatal(e.Start(":" + apiPort))
}
