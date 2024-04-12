package handlers

import (
	"log"
	"net/http"
	"strconv"
	"warranty/api/errs"
	"warranty/database"

	"warranty/api/models"
	"warranty/api/services"

	"github.com/gin-gonic/gin"
)

// Warranty handlers

func GetAllWarranties(c *gin.Context) {
	warranties, err := services.GetAllWarrantiesService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, warranties)
}

func GetWarrantyByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty ID"})
		return
	}

	warranty, err := services.GetWarrantyByIDService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	if warranty == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warranty not found"})
		return
	}

	c.JSON(http.StatusOK, warranty)
}

func CreateWarranty(c *gin.Context) {
	var warranty models.Warranty
	if err := c.BindJSON(&warranty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty data"})
		return
	}

	newWarranty, err := services.CreateWarrantyService(c.Request.Context(), database.GetDB(), &warranty)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newWarranty)
}

func UpdateWarranty(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty ID"})
		return
	}

	var warranty models.Warranty
	if err := c.BindJSON(&warranty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty data"})
		return
	}
	warranty.ID = int64(uint(id))

	updatedWarranty, err := services.UpdateWarrantyService(c.Request.Context(), database.GetDB(), &warranty)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedWarranty)
}

func DeleteWarranty(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty ID"})
		return
	}

	err = services.DeleteWarrantyService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Claim handlers (assuming claims are a sub-resource of warranties)

func GetAllClaimsForWarranty(c *gin.Context) {
	warrantyIDStr := c.Param("warrantyID")
	warrantyID, err := strconv.ParseUint(warrantyIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty ID"})
		return
	}

	claims, err := services.GetAllClaimsForWarrantyService(c.Request.Context(), database.GetDB(), uint(warrantyID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, claims)
}

func handleError(c *gin.Context, err error) {
	// Log the error for further investigation
	log.Printf("Error: %v\n", err)

	// Handle different error types and set appropriate response status code and error message
	switch err {
	case errs.ErrUserNotFound, errs.ErrWarrantyNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Use custom error message
	case errs.ErrUsernameOrEmailExists, errs.ErrInvalidEmailOrPass, errs.ErrInvalidPhone, errs.ErrInvalidUsernameOrPassword:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Use custom error message for validation errors
	default:
		// Handle other unexpected errors (consider returning internal server error for sensitive data)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
