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

// Claim Audit handlers

func GetAllClaimAudits(c *gin.Context) {
	audits, err := services.GetAllClaimAuditsService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, audits)
}

func GetClaimAuditByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid claim audit ID"})
		return
	}

	audit, err := services.GetClaimAuditByIDService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrClaimAuditNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Claim audit not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, audit)
}

func CreateClaimAudit(c *gin.Context) {
	var audit models.ClaimAudit
	if err := c.BindJSON(&audit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid claim audit data"})
		return
	}

	newAudit, err := services.CreateClaimAuditService(c.Request.Context(), database.GetDB(), &audit)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newAudit)
}

func UpdateClaimAudit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid claim audit ID"})
		return
	}

	var audit models.ClaimAudit
	if err := c.BindJSON(&audit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid claim audit data"})
		return
	}
	audit.ID = int64(uint(id)) // Assuming ID field is of type uint

	updatedAudit, err := services.UpdateClaimAuditService(c.Request.Context(), database.GetDB(), &audit)
	if err != nil {
		if errors.Is(err, errs.ErrClaimAuditNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Claim audit not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedAudit)
}

func DeleteClaimAudit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid claim audit ID"})
		return
	}

	err = services.DeleteClaimAuditService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrClaimAuditNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Claim audit not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
