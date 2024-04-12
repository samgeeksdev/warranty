package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"warranty/api/errs"
	"warranty/api/models"
	"warranty/api/services"
	"warranty/database"

	"github.com/gin-gonic/gin"
)

// Product handlers

func GetProducts(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	filter := models.ProductFilter{}

	if categoryIDStr, ok := queryParams["category_id"]; ok && len(categoryIDStr) > 0 {
		categoryID, err := strconv.ParseUint(categoryIDStr[0], 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}
		cID := uint(categoryID)
		filter.CategoryID = &cID
	}

	// You can add support for other filter parameters here (e.g., name, price range)

	products, err := services.GetProductsService(c.Request.Context(), database.GetDB(), &filter)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := services.GetProductService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	newProduct, err := services.CreateProductService(c.Request.Context(), database.GetDB(), &product)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}
	product.ID = int64(uint(id)) // Assuming ID field is of type uint

	updatedProduct, err := services.UpdateProductService(c.Request.Context(), database.GetDB(), &product)
	if err != nil {
		if errors.Is(err, errs.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = services.DeleteProductService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
