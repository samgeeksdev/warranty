package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"warranty/api/errs"
	"warranty/api/models"
	"warranty/api/services"
	"warranty/database"
)

func GetAllWarrantyAudits(c *gin.Context) {
	audits, err := services.GetAllWarrantiesAuditsService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, audits)
}

func GetWarrantyAuditByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty audit ID"})
		return
	}

	audit, err := services.GetWarrantyAuditByIDService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		// Handle error based on the type of error from services.GetWarrantyAuditByIDService
		if errors.Is(err, errs.ErrWarrantyAuditNotFound) { // Use errors.Is for type checking
			c.JSON(http.StatusNotFound, gin.H{"error": "Warranty audit not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, audit)
}

func CreateWarrantyAudit(c *gin.Context) {
	var audit models.WarrantyAudit
	if err := c.BindJSON(&audit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty audit data"})
		return
	}

	newAudit, err := services.CreateWarrantyAuditService(c.Request.Context(), database.GetDB(), &audit)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newAudit)
}

func UpdateWarrantyAudit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty audit ID"})
		return
	}

	var audit models.WarrantyAudit
	if err := c.BindJSON(&audit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty audit data"})
		return
	}
	audit.ID = int64(uint(id)) // Assuming ID field is of type uint

	updatedAudit, err := services.UpdateWarrantyAuditService(c.Request.Context(), database.GetDB(), &audit)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedAudit)
}

func DeleteWarrantyAudit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warranty audit ID"})
		return
	}

	err = services.DeleteWarrantyAuditService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
