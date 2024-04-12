package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserRoleNotFound    = errors.New("user role not found")
	ErrMissingUserRoleData = errors.New("missing user role data")
	// Add other potential errors related to user roles (e.g., ErrDuplicateUserRole)
)

// GetUserRolesService retrieves roles assigned to a user by user ID from the database.
func GetUserRolesService(ctx context.Context, db *gorm.DB, userID uint) ([]models.UserRole, error) {
	var userRoles []models.UserRole
	result := db.WithContext(ctx).Preload("Role").Where("user_id = ?", userID).Find(&userRoles)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to retrieve user roles: %w", err)
	}

	return userRoles, nil
}

// AssignUserRoleService assigns a role to a user in the database.
func AssignUserRoleService(ctx context.Context, db *gorm.DB, userID uint, userRole *models.UserRole) (*models.UserRole, error) {
	if userRole == nil || userRole.UserID == 0 || userRole.RoleID == 0 {
		return nil, ErrMissingUserRoleData
	}

	// Implement logic to check if user exists before assigning a role
	// if err := ValidateUser(userID, db); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&userRole)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to assign user role: %w", err)
	}

	return userRole, nil
}

// RemoveUserRoleService removes a role assigned to a user by user ID and role ID from the database.
func RemoveUserRoleService(ctx context.Context, db *gorm.DB, userID, roleID uint) error {
	result := db.WithContext(ctx).Delete(&models.UserRole{}, map[string]interface{}{"user_id": userID, "role_id": roleID})
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserRoleNotFound
		}
		return fmt.Errorf("failed to remove user role: %w", err)
	}

	return nil
}
