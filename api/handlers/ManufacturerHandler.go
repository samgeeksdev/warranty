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

// Manufacturer handlers

func GetManufacturers(c *gin.Context) {
	manufacturers, err := services.GetManufacturersService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, manufacturers)
}

func CreateManufacturer(c *gin.Context) {
	var manufacturer models.Manufacturer
	if err := c.BindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid manufacturer data"})
		return
	}

	newManufacturer, err := services.CreateManufacturerService(c.Request.Context(), database.GetDB(), &manufacturer)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newManufacturer)
}

func GetManufacturerByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid manufacturer ID"})
		return
	}

	manufacturer, err := services.GetManufacturerByIDService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrManufacturerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Manufacturer not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, manufacturer)
}

func UpdateManufacturer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid manufacturer ID"})
		return
	}

	var manufacturer models.Manufacturer
	if err := c.BindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid manufacturer data"})
		return
	}
	manufacturer.ID = int64(uint(id)) // Assuming ID field is of type uint

	updatedManufacturer, err := services.UpdateManufacturerService(c.Request.Context(), database.GetDB(), &manufacturer)
	if err != nil {
		if errors.Is(err, errs.ErrManufacturerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Manufacturer not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedManufacturer)
}

func DeleteManufacturer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid manufacturer ID"})
		return
	}

	err = services.DeleteManufacturerService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrManufacturerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Manufacturer not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
