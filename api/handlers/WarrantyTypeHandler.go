package handlers

import (
	"context"
	"net/http"
	"strconv"
	"warranty/api/errs"
	"warranty/database"

	"github.com/gin-gonic/gin"

	"warranty/api/models"
	"warranty/api/services" // Now imported explicitly
)

// GetWarrantyTypes handles GET request to fetch all warranty types
func GetWarrantyTypes(c *gin.Context) {
	ctx := context.Background()
	warrantyTypes, err := services.GetAllWarrantyTypes(database.GetDB(), ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, warrantyTypes)
}

// CreateWarrantyType handles POST request to create a new warranty type
func CreateWarrantyType(c *gin.Context) {
	var newWarrantyType models.WarrantyType
	// Decode request body into the newWarrantyType struct
	if err := c.BindJSON(&newWarrantyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	createdWarrantyType, err := services.CreateWarrantyType(database.GetDB(), ctx, newWarrantyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdWarrantyType)
}

// GetWarrantyTypeByID handles GET request to fetch a warranty type by ID
func GetWarrantyTypeByID(c *gin.Context) {
	id := c.Param("id")

	warrantyType, err := services.GetWarrantyTypeByID(database.GetDB(), c.Request.Context(), id)
	if err != nil {
		if err == errs.ErrWarrantyTypeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "warranty type not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, warrantyType)
}

// UpdateWarrantyType handles PUT request to update a warranty type
func UpdateWarrantyType(c *gin.Context) {
	var updatedWarrantyType models.WarrantyType

	// Decode request body into the updatedWarrantyType struct
	if err := c.BindJSON(&updatedWarrantyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call UpdateWarrantyType service function
	updatedType, err := services.UpdateWarrantyType(database.GetDB(), c.Request.Context(), updatedWarrantyType)
	if err != nil {
		if err == errs.ErrWarrantyTypeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "warranty type not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedType)
}

// DeleteWarrantyType handles DELETE request to delete a warranty type
func DeleteWarrantyType(c *gin.Context) {
	id := c.Param("id")

	// Parse ID to integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	// Call DeleteWarrantyType service function
	err = services.DeleteWarrantyType(database.GetDB(), c.Request.Context(), strconv.Itoa(idInt))
	if err != nil {
		if err == errs.ErrWarrantyTypeNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "warranty type not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

//func parseInt(str string) int64 {
//	id, err := fmt.strconv.ParseInt(str, 10, 64)
//	if err != nil {
//		panic(err) // Handle ID parsing error more gracefully in production
//	}
//	return id
//}
