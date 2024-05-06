package product

import (
 "net/http"
 "github.com/nuchit2019/product-res-api/config"

 "github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
 db := config.DB()
 var products []Product

 if err := db.Find(&products).Error; err != nil {
  res := Response{
   Status:  http.StatusNotFound,
   Message: "Failed to retrieve products...Err:" + err.Error(),
  }
  return c.JSON(http.StatusInternalServerError, res)
 }

 res := Response{
  Status:  http.StatusOK,
  Message: "successfully",
  Data:    products,
 }

 return c.JSON(http.StatusOK, res)

}

func GetProduct(c echo.Context) error {
 id := c.Param("id")
 db := config.DB()
 var product Product
 if err := db.First(&product, id).Error; err != nil {
  res := Response{
   Status:  http.StatusNotFound,
   Message: "Failed to retrieve product id:=" + id + "...Err:" + err.Error(),
  }
  return c.JSON(http.StatusNotFound, res)
 }

 res := Response{
  Status:  http.StatusOK,
  Message: "successfully",
  Data:    product,
 }

 return c.JSON(http.StatusOK, res)
}

func CreateProduct(c echo.Context) error {
 product := new(Product)

 if err := c.Bind(product); err != nil {
  res := Response{
   Status:  http.StatusBadRequest,
   Message: "Failed to create product...Err:" + err.Error(),
  }
  return c.JSON(http.StatusBadRequest, res)
 }

 db := config.DB()

 if err := db.Create(product).Error; err != nil {
  res := Response{
   Status:  http.StatusNotFound,
   Message: "Failed to create product...Err:" + err.Error(),
  }
  return c.JSON(http.StatusInternalServerError, res)
 }

 res := Response{
  Status:  http.StatusCreated,
  Message: "successfully",
  Data:    product,
 }
 return c.JSON(http.StatusCreated, res)

}

func UpdateProduct(c echo.Context) error {
 id := c.Param("id")
 product := new(Product)
 if err := c.Bind(product); err != nil {
  res := Response{
   Status:  http.StatusBadRequest,
   Message: "Failed to update product...Err:" + err.Error(),
  }
  return c.JSON(http.StatusBadRequest, res)
 }

 existingProduct := Product{}
 if err := config.DB().First(&existingProduct, id).Error; err != nil {
  res := Response{
   Status:  http.StatusNotFound,
   Message: "Failed to update product...Err:" + err.Error(),
  }
  return c.JSON(http.StatusNotFound, res)
 }

 if err := config.DB().Model(&Product{}).Where("id =?", id).Updates(product).Error; err != nil {
  res := Response{
   Status:  http.StatusNotFound,
   Message: "Failed to update product...Err:" + err.Error(),
  }
  return c.JSON(http.StatusInternalServerError, res)

 }

 updatedProduct := Product{}
 if err := config.DB().First(&updatedProduct, id).Error; err != nil {
  res := Response{
   Status:  http.StatusNotFound,
   Message: "Failed to update product...Err:" + err.Error(),
  }
  return c.JSON(http.StatusInternalServerError, res)
 }

 res := Response{
  Status:  http.StatusOK,
  Message: "successfully",
  Data:    updatedProduct,
 }
 return c.JSON(http.StatusOK, res)
}

func DeleteProduct(c echo.Context) error {
 id := c.Param("id")
 db := config.DB()

 existingProduct := Product{}
 if err := config.DB().First(&existingProduct, id).Error; err != nil {
  res := Response{
   Status:  http.StatusNotFound,
   Message: "Failed to delete product...Err:" + err.Error(),
  }
  return c.JSON(http.StatusNotFound, res)
 }

 if err := db.Delete(&Product{}, id).Error; err != nil {
  res := Response{
   Status:  http.StatusNotFound,
   Message: "Failed to delete product...Err:" + err.Error(),
  }
  return c.JSON(http.StatusInternalServerError, res)
 }

 res := Response{
  Status:  http.StatusOK,
  Message: "Delete product successfully",
  Data:    nil, 
 }

 return c.JSON(http.StatusOK, res)
}