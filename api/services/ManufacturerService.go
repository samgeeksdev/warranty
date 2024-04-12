package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	ErrManufacturerNotFound    = errors.New("manufacturer not found")
	ErrMissingManufacturerData = errors.New("missing manufacturer data")
	// Add other potential errors related to manufacturers (e.g., ErrDuplicateManufacturer, ErrInvalidManufacturerName)
)

// Services for manufacturers

func GetManufacturersService(ctx context.Context, db *gorm.DB) ([]models.Manufacturer, error) {
	var manufacturers []models.Manufacturer
	result := db.WithContext(ctx).Find(&manufacturers)
	if err := result.Error; err != nil {
		return nil, err
	}

	return manufacturers, nil
}

func GetManufacturerByIDService(ctx context.Context, db *gorm.DB, id uint) (*models.Manufacturer, error) {
	var manufacturer models.Manufacturer
	result := db.WithContext(ctx).Where("id = ?", id).First(&manufacturer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrManufacturerNotFound
		}
		return nil, err
	}

	return &manufacturer, nil
}

// CreateManufacturerService creates a new manufacturer in the database
func CreateManufacturerService(ctx context.Context, db *gorm.DB, manufacturer *models.Manufacturer) (*models.Manufacturer, error) {
	if manufacturer == nil {
		return nil, ErrMissingManufacturerData
	}

	// Implement validation logic for manufacturer data before creation (e.g., name format)
	// if err := ValidateManufacturer(manufacturer); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&manufacturer)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create manufacturer: %w", err)
	}

	return manufacturer, nil
}

// UpdateManufacturerService updates an existing manufacturer in the database
func UpdateManufacturerService(ctx context.Context, db *gorm.DB, manufacturer *models.Manufacturer) (*models.Manufacturer, error) {
	if manufacturer == nil || manufacturer.ID == 0 {
		return nil, ErrMissingManufacturerData
	}

	// Implement validation logic for manufacturer data before update (e.g., name format)
	// if err := ValidateManufacturer(manufacturer); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Save(&manufacturer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrManufacturerNotFound
		}
		return nil, fmt.Errorf("failed to update manufacturer: %w", err)
	}

	return manufacturer, nil
}

// DeleteManufacturerService deletes a manufacturer from the database by ID
func DeleteManufacturerService(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&models.Manufacturer{}, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrManufacturerNotFound
		}
		return fmt.Errorf("failed to delete manufacturer: %w", err)
	}

	return nil
}
