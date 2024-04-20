package controller

import (
	"net/http"
	"product-res-api/config"
	"product-res-api/model"

	"github.com/labstack/echo/v4"
)

// CreateProduct creates a new product based on the request data
func GetProducts(c echo.Context) error {
	// Get the database instance
	db := config.DB()

	// Create a slice to store products retrieved from the database
	var products []model.Product

	// Query the database to fetch all products
	if err := db.Find(&products).Error; err != nil {
		// If an error occurs during the database query, return an internal server error response

		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to retrieve products...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// Create a response map to format the JSON response
	res := model.Response{
		Status:  http.StatusOK,
		Message: "successfully",
		Data:    products,
	}

	// Return a JSON response with the HTTP status OK (200) and the formatted response map
	return c.JSON(http.StatusOK, res)

}

// GetProduct retrieves a product by its ID from the database.
func GetProduct(c echo.Context) error {
	// Extract product ID from request parameters
	id := c.Param("id")

	// Get the database instance
	db := config.DB()

	// Initialize a variable to store the retrieved product
	var product model.Product

	// Query the database to find the product with the given ID
	if err := db.First(&product, id).Error; err != nil {
		// If the product is not found, return a not found response
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to retrieve product id:=" + id + "...Err:" + err.Error(),
		}
		return c.JSON(http.StatusNotFound, res)
	}

	// Create a success response with the retrieved product data
	res := model.Response{
		Status:  http.StatusOK,
		Message: "successfully",
		Data:    product,
	}

	// Return a JSON response with the success response
	return c.JSON(http.StatusOK, res)
}

// CreateProduct creates a new product based on the request data
func CreateProduct(c echo.Context) error {
	// Create a new product instance
	product := new(model.Product)

	// Bind the request body to the product struct
	if err := c.Bind(product); err != nil {
		// If there's an error in binding the request body, return a bad request response
		res := model.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to create product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// Get the database instance
	db := config.DB()

	// Create the product in the database
	if err := db.Create(product).Error; err != nil {
		// If there's an error creating the product, return an internal server error response
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to create product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// If creation is successful, return a success response with the created product data
	res := model.Response{
		Status:  http.StatusCreated,
		Message: "successfully",
		Data:    product,
	}
	return c.JSON(http.StatusCreated, res)

}

// UpdateProduct updates a product by its ID from the database.
func UpdateProduct(c echo.Context) error {

	// Extract product ID from request parameters
	id := c.Param("id")

	// Create a new product instance to hold updated data
	product := new(model.Product)

	// Bind the request body to the product struct
	if err := c.Bind(product); err != nil {
		// If there's an error in binding the request body, return a bad request response
		res := model.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to update product...Err:" + err.Error(),
		}

		return c.JSON(http.StatusBadRequest, res)
	}

	// Check if the product exists in the database
	existingProduct := model.Product{}
	if err := config.DB().First(&existingProduct, id).Error; err != nil {
		// If the product doesn't exist, return a not found response
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to update product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusNotFound, res)
	}

	// Update the product in the database
	if err := config.DB().Model(&model.Product{}).Where("id =?", id).Updates(product).Error; err != nil {
		// If there's an error updating the product, return an internal server error response
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to update product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)

	}

	// Read the updated product from the database
	updatedProduct := model.Product{}
	if err := config.DB().First(&updatedProduct, id).Error; err != nil {
		// If there's an error reading the updated product, return an internal server error response
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to update product...Err:" + err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, res)
	}

	// Create a success response with the updated product data
	res := model.Response{
		Status:  http.StatusOK,
		Message: "successfully",
		Data:    updatedProduct,
	}

	// Return a JSON response with the success response
	return c.JSON(http.StatusOK, res)
}

// DeleteProduct deletes a product by its ID from the database.
func DeleteProduct(c echo.Context) error {
	// Extract product ID from request parameters
	id := c.Param("id")

	// Get the database instance
	db := config.DB()

	// Check if the product exists in the database
	existingProduct := model.Product{}
	if err := config.DB().First(&existingProduct, id).Error; err != nil {
		// If the product doesn't exist, return a not found response
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to delete product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusNotFound, res)
	}

	// Delete the product from the database
	if err := db.Delete(&model.Product{}, id).Error; err != nil {
		// If there's an error deleting the product, return an internal server error response
		res := model.Response{
			Status:  http.StatusNotFound,
			Message: "Failed to delete product...Err:" + err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// If deletion is successful, return a success response
	res := model.Response{
		Status:  http.StatusOK,
		Message: "Delete product successfully",
		Data:    nil,// No data to return after deletion
	}

	// Return a JSON response with the success response
	return c.JSON(http.StatusOK, res)
}
