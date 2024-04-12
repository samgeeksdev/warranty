package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	ErrPermissionNotFound    = errors.New("permission not found")
	ErrMissingPermissionData = errors.New("missing permission data")
	// Add other potential errors related to permissions (e.g., ErrDuplicatePermission, ErrInvalidPermissionName)
)

// Services for permissions

// GetPermissionsService retrieves all permissions from the database.
func GetPermissionsService(ctx context.Context, db *gorm.DB) ([]models.Permission, error) {
	var permissions []models.Permission
	result := db.WithContext(ctx).Find(&permissions)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve permissions: %w", err)
	}

	return permissions, nil
}

// CreatePermissionService creates a new permission in the database.
func CreatePermissionService(ctx context.Context, db *gorm.DB, permission *models.Permission) (*models.Permission, error) {
	if permission == nil {
		return nil, ErrMissingPermissionData
	}

	// Implement validation logic for permission data before creation (e.g., name format)
	// if err := ValidatePermission(permission); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&permission)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create permission: %w", err)
	}

	return permission, nil
}

// GetPermissionByIDService retrieves a permission by ID from the database.
func GetPermissionByIDService(ctx context.Context, db *gorm.DB, id uint) (*models.Permission, error) {
	var permission models.Permission
	result := db.WithContext(ctx).Where("id = ?", id).First(&permission)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPermissionNotFound
		}
		return nil, fmt.Errorf("failed to retrieve permission by ID: %w", err)
	}

	return &permission, nil
}

// UpdatePermissionService updates an existing permission in the database.
func UpdatePermissionService(ctx context.Context, db *gorm.DB, permission *models.Permission) (*models.Permission, error) {
	if permission == nil || permission.ID == 0 {
		return nil, ErrMissingPermissionData
	}

	// Implement validation logic for permission data before update (e.g., name format)
	// if err := ValidatePermission(permission); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Save(&permission)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPermissionNotFound
		}
		return nil, fmt.Errorf("failed to update permission: %w", err)
	}

	return permission, nil
}

// DeletePermissionService deletes a permission by ID from the database.
func DeletePermissionService(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&models.Permission{}, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPermissionNotFound
		}
		return fmt.Errorf("failed to delete permission: %w", err)
	}

	return nil
}
