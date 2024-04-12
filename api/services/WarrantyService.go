package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"warranty/api/errs"
	"warranty/api/models"
)

func GetAllWarrantiesService(ctx context.Context, db *gorm.DB) ([]models.Warranty, error) {
	var warranty []models.Warranty
	result := db.WithContext(ctx).Find(&warranty)

	if result.Error != nil {
		return nil, result.Error
	}
	return warranty, nil

}

func GetWarrantyByIDService(ctx context.Context, db *gorm.DB, id uint) (*models.Warranty, error) {
	// Find the warranty record in the database by ID
	var warranty models.Warranty
	result := db.WithContext(ctx).First(&warranty, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrWarrantyNotFound // Use custom error for clarity
		}
		return nil, err
	}
	return &warranty, nil
}

func CreateWarrantyService(ctx context.Context, db *gorm.DB, warranty *models.Warranty) (*models.Warranty, error) {
	// Validate warranty data before creation (implement validation logic)
	// if err := ValidateWarranty(warranty); err != nil {
	//   return nil, err
	// }

	// Create a new warranty record in the database
	result := db.WithContext(ctx).Create(&warranty)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create warranty: %w", err) // Wrap general errors
	}

	return warranty, nil
}

func UpdateWarrantyService(ctx context.Context, db *gorm.DB, warranty *models.Warranty) (*models.Warranty, error) {
	// Validate warranty data before update (implement validation logic)
	// if err := ValidateWarranty(warranty); err != nil {
	//   return nil, err
	// }

	// Update the existing warranty record in the database
	result := db.WithContext(ctx).Save(&warranty)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrWarrantyNotFound // Use custom error for clarity
		}
		return nil, err
	}

	return warranty, nil
}

func DeleteWarrantyService(ctx context.Context, db *gorm.DB, id uint) error {
	// Check for associated claims referencing the warranty
	var claims []models.Claim
	result := db.WithContext(ctx).Where("warranty_id = ?", id).Find(&claims)
	if result.Error != nil {
		return errs.ErrWarrantyNotFound // Handle error finding claims
	}

	// Option 1: Delete associated claims first (if allowed)
	if len(claims) > 0 {
		// Implement logic to delete claims (e.g., a batch delete or loop)
		// ...
	}

	// Delete the warranty
	result = db.WithContext(ctx).Delete(&models.Warranty{}, id)
	if err := result.Error; err != nil {
		return err // Handle error deleting warranty
	}

	return nil // Warranty deleted successfully
}
func GetAllClaimsForWarrantyService(ctx context.Context, db *gorm.DB, warrantyID uint) ([]models.Claim, error) {
	claims, err := GetAllClaimsForWarranty(ctx, db, warrantyID)
	if err != nil {
		return nil, err
	}
	// Dereference the pointer to return the actual slice
	return *claims, nil
}

func GetAllClaimsForWarranty(ctx context.Context, db *gorm.DB, warrantyID uint) (*[]models.Claim, error) {
	var claims []models.Claim
	result := db.WithContext(ctx).Where("warranty_id = ?", warrantyID).Find(&claims)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// No claims found, not an error
			return &claims, nil
		}
		return nil, result.Error
	}

	return &claims, nil
}

// GetClaimByID retrieves a specific claim by its ID within a warranty
