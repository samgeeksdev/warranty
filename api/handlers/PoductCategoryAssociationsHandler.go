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

// Product Category Association handlers

func GetAllProductCategoryAssociations(c *gin.Context) {
	associations, err := services.GetAllProductCategoryAssociationsService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, associations)
}

func GetProductCategoryAssociation(c *gin.Context) {
	productIDStr := c.Param("product_id")
	categoryIDStr := c.Param("category_id")

	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	association, err := services.GetProductCategoryAssociationService(c.Request.Context(), database.GetDB(), uint(productID), uint(categoryID))
	if err != nil {
		if errors.Is(err, errs.ErrProductCategoryAssociationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product category association not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, association)
}

func CreateProductCategoryAssociation(c *gin.Context) {
	var association models.ProductCategoryAssociation
	if err := c.BindJSON(&association); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product category association data"})
		return
	}

	newAssociation, err := services.CreateProductCategoryAssociationService(c.Request.Context(), database.GetDB(), &association)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newAssociation)
}

func DeleteProductCategoryAssociation(c *gin.Context) {
	productIDStr := c.Param("product_id")
	categoryIDStr := c.Param("category_id")

	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	err = services.DeleteProductCategoryAssociationService(c.Request.Context(), database.GetDB(), uint(productID), uint(categoryID))
	if err != nil {
		if errors.Is(err, errs.ErrProductCategoryAssociationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product category association not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
