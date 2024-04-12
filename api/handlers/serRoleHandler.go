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

func GetUserRoles(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userRoles, err := services.GetUserRolesService(c.Request.Context(), database.GetDB(), uint(userID))
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, userRoles)
}

func AssignUserRole(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userRole := models.UserRole{}
	if err := c.BindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user role data"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	assignedUserRole, err := services.AssignUserRoleService(c.Request.Context(), database.GetDB(), uint(userID), &userRole)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, assignedUserRole)
}

func RemoveUserRole(c *gin.Context) {
	userIDStr := c.Param("user_id")
	roleIDStr := c.Param("role_id")

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

	err = services.RemoveUserRoleService(c.Request.Context(), database.GetDB(), uint(userID), uint(roleID))
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
