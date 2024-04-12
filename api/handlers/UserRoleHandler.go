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

// User Role handlers

func GetAllUserRoles(c *gin.Context) {
	userRoles, err := services.GetAllUserRolesService(c.Request.Context(), database.GetDB())
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, userRoles)
}

func GetUserRole(c *gin.Context) {
	userIDStr := c.Param("userID")
	roleIDStr := c.Param("roleID")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	roleID, err := strconv.ParseUint(roleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	userRole, err := services.GetUserRoleService(c.Request.Context(), database.GetDB(), uint(userID), uint(roleID))
	if err != nil {
		if errors.Is(err, errs.ErrUserRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User role not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, userRole)
}

func CreateUserRole(c *gin.Context) {
	var userRole models.UserRole
	if err := c.BindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user role data"})
		return
	}

	newUserRole, err := services.CreateUserRoleService(c.Request.Context(), database.GetDB(), &userRole)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newUserRole)
}

func DeleteUserRole(c *gin.Context) {
	userIDStr := c.Param("userID")
	roleIDStr := c.Param("roleID")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	roleID, err := strconv.ParseUint(roleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	err = services.DeleteUserRoleService(c.Request.Context(), database.GetDB(), uint(userID), uint(roleID))
	if err != nil {
		if errors.Is(err, errs.ErrUserRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User role not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// You can add more handlers for user roles based on your requirements
