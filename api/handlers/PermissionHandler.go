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

// Permission handlers

func GetPermissions(c *gin.Context) {
	permissions, err := services.GetPermissionsService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, permissions)
}

func CreatePermission(c *gin.Context) {
	var permission models.Permission
	if err := c.BindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission data"})
		return
	}

	newPermission, err := services.CreatePermissionService(c.Request.Context(), database.GetDB(), &permission)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newPermission)
}

func GetPermissionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}

	permission, err := services.GetPermissionByIDService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrPermissionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, permission)
}

func UpdatePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}

	var permission models.Permission
	if err := c.BindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission data"})
		return
	}
	permission.ID = int64(uint(id)) // Assuming ID field is of type uint

	updatedPermission, err := services.UpdatePermissionService(c.Request.Context(), database.GetDB(), &permission)
	if err != nil {
		if errors.Is(err, errs.ErrPermissionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedPermission)
}

func DeletePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid permission ID"})
		return
	}

	err = services.DeletePermissionService(c.Request.Context(), database.GetDB(), uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrPermissionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
