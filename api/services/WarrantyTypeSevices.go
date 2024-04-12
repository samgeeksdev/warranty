package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"warranty/api/errs"

	"warranty/api/models"
)

// GetAllWarrantyTypes retrieves all warranty types from the database.
func GetAllWarrantyTypes(db *gorm.DB, ctx context.Context) ([]models.WarrantyType, error) {
	var warrantyTypes []models.WarrantyType
	result := db.WithContext(ctx).Find(&warrantyTypes)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve warranty types: %w", err)
	}

	return warrantyTypes, nil
}

// CreateWarrantyType creates a new warranty type in the database.
func CreateWarrantyType(db *gorm.DB, ctx context.Context, warrantyType models.WarrantyType) (*models.WarrantyType, error) {
	// Implement validation logic for warranty type data before creation (e.g., required fields)
	// if err := ValidateWarrantyType(warrantyType); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&warrantyType)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create warranty type: %w", err)
	}

	return &warrantyType, nil
}

// GetWarrantyTypeByID retrieves a specific warranty type by ID from the database.
func GetWarrantyTypeByID(db *gorm.DB, ctx context.Context, id string) (*models.WarrantyType, error) {
	var warrantyType models.WarrantyType
	result := db.WithContext(ctx).Where("id = ?", id).First(&warrantyType)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrWarrantyTypeNotFound
		}
		return nil, fmt.Errorf("failed to retrieve warranty type: %w", err)
	}

	return &warrantyType, nil
}

// UpdateWarrantyType updates an existing warranty type in the database.
func UpdateWarrantyType(db *gorm.DB, ctx context.Context, warrantyType models.WarrantyType) (*models.WarrantyType, error) {
	if warrantyType.ID == 0 {
		return nil, errors.New("missing warranty type ID")
	}

	// Implement validation logic for warranty type data before update (e.g., required fields)
	// if err := ValidateWarrantyType(warrantyType); err != nil {
	//   return nil, err
	// }

	// Convert warrantyType.ID to type int (assuming ID field is int64)
	result := db.WithContext(ctx).Where("id = ?", int(warrantyType.ID)).Save(&warrantyType)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrWarrantyTypeNotFound
		}
		return nil, fmt.Errorf("failed to update warranty type: %w", err)
	}

	return &warrantyType, nil
}

// DeleteWarrantyType deletes a warranty type by ID from the database.
func DeleteWarrantyType(db *gorm.DB, ctx context.Context, id string) error {
	result := db.WithContext(ctx).Delete(&models.WarrantyType{}, map[string]interface{}{"id": id})
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrWarrantyTypeNotFound
		}
		return fmt.Errorf("failed to delete warranty type: %w", err)
	}

	return nil
}
