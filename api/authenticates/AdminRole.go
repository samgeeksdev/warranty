package authenticates

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"warranty/api/models"
)

func FindUserRoleByID(ctx context.Context, db *gorm.DB, userID uint) (*models.UserRole, error) {
	var userRole models.UserRole

	// Find the user role record in the database by user_id
	result := db.WithContext(ctx).Where("user_id = ?", userID).First(&userRole)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user role not found for user ID %d", userID)
		}
		return nil, fmt.Errorf("failed to find user role: %w", err)
	}

	return &userRole, nil
}

func IsAdmin(db *gorm.DB, user_id uint) (bool, error) {
	var Roles models.Role
	ctx := context.Background()

	Role, err := FindUserRoleByID(ctx, db, user_id)
	if err != nil {
		return false, fmt.Errorf("failed to find user role: %v", err)
	}

	if err := db.WithContext(ctx).Unscoped().Where("id = ?", Role.RoleID).First(&Roles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("role not found for role ID %d", Role.RoleID)
		}
		return false, fmt.Errorf("failed to retrieve user role: %v", err)
	}

	if Roles.Name != "admin" {
		return false, fmt.Errorf("user is not an admin")
	}
	return true, nil
}

func IsEditor(db *gorm.DB, user_id uint) (bool, error) {
	var Roles models.Role
	ctx := context.Background()

	Role, err := FindUserRoleByID(ctx, db, user_id)
	if err != nil {
		return false, fmt.Errorf("failed to find user role: %v", err)
	}

	if err := db.WithContext(ctx).Unscoped().Where("id = ?", Role.RoleID).First(&Roles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("role not found for role ID %d", Role.RoleID)
		}
		return false, fmt.Errorf("failed to retrieve user role: %v", err)
	}

	if Roles.Name != "Editor" {
		return false, fmt.Errorf("user is not an admin")
	}
	return true, nil
}

func IsUser(db *gorm.DB, user_id uint) (bool, error) {
	var Roles models.Role
	ctx := context.Background()

	Role, err := FindUserRoleByID(ctx, db, user_id)
	if err != nil {
		return false, fmt.Errorf("failed to find user role: %v", err)
	}

	if err := db.WithContext(ctx).Unscoped().Where("id = ?", Role.RoleID).First(&Roles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("role not found for role ID %d", Role.RoleID)
		}
		return false, fmt.Errorf("failed to retrieve user role: %v", err)
	}

	if Roles.Name != "user" {
		return false, fmt.Errorf("user is not an admin")
	}
	return true, nil
}
func Authorize(db *gorm.DB, role string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Extract user ID from Gin context (replace with your specific logic)
		userID, err := ExtractUserID(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var authorized bool
		switch role {
		case "admin":
			authorized, err = IsAdmin(db, userID)
		case "editor":
			authorized, err = IsEditor(db, userID)
		case "user":
			authorized, err = IsUser(db, userID)
		default:
			err = fmt.Errorf("unknown role: %s", role)
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !authorized {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// Call the next handler in the chain if authorized
		c.Next()
	}
}

// Assuming you have a way to extract user ID from context (e.g., from JWT claims)
func ExtractUserID(c *gin.Context) (uint, error) {
	// Replace with your logic for user extraction from Gin context
	userIDInt, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("user ID not found in context")
	}

	userID, ok := userIDInt.(uint)
	if !ok {
		return 0, errors.New("invalid user ID type")
	}

	return userID, nil
}

func AdminOnly(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Extract user ID from Gin context (replace with your specific logic)
		userID, err := ExtractUserID(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		isAdmin, err := IsAdmin(db, userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !isAdmin {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// Call the next handler in the chain if authorized
		c.Next()
	}
}
