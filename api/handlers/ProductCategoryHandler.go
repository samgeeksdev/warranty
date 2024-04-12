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

// Product Category handlers

func GetProductCategories(c *gin.Context) {
	productCategories, err := services.GetAllProductCategoriesService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, productCategories)
}

func GetProductCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product category ID"})
		return
	}

	productCategory, err := services.GetProductCategoryByIDService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrProductCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product category not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, productCategory)
}

func CreateProductCategory(c *gin.Context) {
	var productCategory models.ProductCategory
	if err := c.BindJSON(&productCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product category data"})
		return
	}

	newProductCategory, err := services.CreateProductCategoryService(c.Request.Context(), database.GetDB(), &productCategory)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newProductCategory)
}

func UpdateProductCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product category ID"})
		return
	}

	var productCategory models.ProductCategory
	if err := c.BindJSON(&productCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product category data"})
		return
	}
	productCategory.ID = int64(uint(id)) // Assuming ID field is of type uint

	updatedProductCategory, err := services.UpdateProductCategoryService(c.Request.Context(), database.GetDB(), &productCategory)
	if err != nil {
		if errors.Is(err, errs.ErrProductCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product category not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedProductCategory)
}

func DeleteProductCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product category ID"})
		return
	}

	err = services.DeleteProductCategoryService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrProductCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product category not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// You can add more handlers for product categories based on your requirements (e.g., get product categories with subcategories)
