package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

// GetAllUserRolesService retrieves all user roles from the database.
func GetAllUserRolesService(ctx context.Context, db *gorm.DB) ([]models.UserRole, error) {
	var userRoles []models.UserRole
	result := db.WithContext(ctx).Preload("User").Preload("Role").Find(&userRoles)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all user roles: %w", err)
	}

	return userRoles, nil
}

// GetUserRoleService retrieves a specific user role by user ID and role ID from the database.
func GetUserRoleService(ctx context.Context, db *gorm.DB, userID, roleID uint) (*models.UserRole, error) {
	var userRole models.UserRole
	result := db.WithContext(ctx).Preload("User").Preload("Role").Where("user_id = ? AND role_id = ?", userID, roleID).First(&userRole)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserRoleNotFound
		}
		return nil, fmt.Errorf("failed to retrieve user role: %w", err)
	}

	return &userRole, nil
}

// CreateUserRoleService creates a new user role in the database.
func CreateUserRoleService(ctx context.Context, db *gorm.DB, userRole *models.UserRole) (*models.UserRole, error) {
	if userRole == nil || userRole.UserID == 0 || userRole.RoleID == 0 {
		return nil, ErrMissingUserRoleData
	}

	// Implement logic to check if user and role exist before assigning
	// if err := ValidateUser(userRole.UserID, db); err != nil {
	//   return nil, err
	// }
	// if err := ValidateRole(userRole.RoleID, db); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&userRole)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create user role: %w", err)
	}

	return userRole, nil
}

// DeleteUserRoleService deletes a user role by user ID and role ID from the database.
func DeleteUserRoleService(ctx context.Context, db *gorm.DB, userID, roleID uint) error {
	result := db.WithContext(ctx).Delete(&models.UserRole{}, map[string]interface{}{"user_id": userID, "role_id": roleID})
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserRoleNotFound
		}
		return fmt.Errorf("failed to delete user role: %w", err)
	}

	return nil
}
