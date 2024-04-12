package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

// These errors are specific to roles and their operations
var (
	ErrRoleNotFound    = errors.New("role not found")
	ErrMissingRoleData = errors.New("missing role data")
	ErrDuplicateRole   = errors.New("duplicate role name found") // Added for potential uniqueness constraint violation
)

// GetAllRolesService retrieves all roles from the database.
func GetAllRolesService(ctx context.Context, db *gorm.DB) ([]models.Role, error) {
	var roles []models.Role
	result := db.WithContext(ctx).Find(&roles)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve roles: %w", err)
	}

	return roles, nil
}

// GetRoleByIDService retrieves a role by its ID from the database.
func GetRoleByIDService(ctx context.Context, db *gorm.DB, id uint) (*models.Role, error) {
	var role models.Role
	result := db.WithContext(ctx).Where("id = ?", id).First(&role)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRoleNotFound
		}
		return nil, fmt.Errorf("failed to retrieve role: %w", err)
	}

	return &role, nil
}

// CreateRoleService creates a new role in the database.
func CreateRoleService(ctx context.Context, db *gorm.DB, role *models.Role) (*models.Role, error) {
	if role == nil {
		return nil, ErrMissingRoleData
	}

	// Implement validation logic for role data before creation (e.g., name format, uniqueness)
	// if err := ValidateRoleName(role.Name, db); err != nil {
	//   return nil, err // Handle potential validation errors (e.g., ErrDuplicateRole)
	// }

	result := db.WithContext(ctx).Create(&role)
	if err := result.Error; err != nil {
		// Handle potential uniqueness constraint violation
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDuplicateRole
		}
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	return role, nil
}

// UpdateRoleService updates an existing role in the database.
func UpdateRoleService(ctx context.Context, db *gorm.DB, role *models.Role) (*models.Role, error) {
	if role == nil || role.ID == 0 {
		return nil, ErrMissingRoleData
	}

	// Implement validation logic for role data before update (e.g., name format, uniqueness on update)
	// if err := ValidateRoleNameOnUpdate(role.ID, role.Name, db); err != nil {
	//   return nil, err // Handle potential validation errors (e.g., ErrDuplicateRole)
	// }

	result := db.WithContext(ctx).Save(&role)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRoleNotFound
		}
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	return role, nil
}

// DeleteRoleService deletes a role by its ID from the database.
func DeleteRoleService(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&models.Role{}, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrRoleNotFound
		}
		return fmt.Errorf("failed to delete role: %w", err)
	}

	return nil
}
