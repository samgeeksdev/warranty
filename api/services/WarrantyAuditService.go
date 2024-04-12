package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	ErrWarrantyAuditNotFound    = errors.New("warranty audit not found")
	ErrMissingWarrantyAuditData = errors.New("missing warranty audit data")
	// Add other potential errors related to warranty audits (e.g., ErrDuplicateWarrantyAudit)
)

// GetAllWarrantiesAuditsService retrieves all warranty audits from the database.
func GetAllWarrantiesAuditsService(ctx context.Context, db *gorm.DB) ([]models.WarrantyAudit, error) {
	var audits []models.WarrantyAudit
	result := db.WithContext(ctx).Find(&audits)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve warranty audits: %w", err)
	}

	return audits, nil
}

// GetWarrantyAuditByIDService retrieves a specific warranty audit by ID from the database.
func GetWarrantyAuditByIDService(ctx context.Context, db *gorm.DB, id uint) (*models.WarrantyAudit, error) {
	var audit models.WarrantyAudit
	result := db.WithContext(ctx).Where("id = ?", id).First(&audit)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWarrantyAuditNotFound
		}
		return nil, fmt.Errorf("failed to retrieve warranty audit: %w", err)
	}

	return &audit, nil
}

// CreateWarrantyAuditService creates a new warranty audit in the database.
func CreateWarrantyAuditService(ctx context.Context, db *gorm.DB, audit *models.WarrantyAudit) (*models.WarrantyAudit, error) {
	if audit == nil {
		return nil, ErrMissingWarrantyAuditData
	}

	// Implement validation logic for warranty audit data before creation
	// (e.g., check for required fields)
	// if err := ValidateWarrantyAudit(audit); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&audit)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create warranty audit: %w", err)
	}

	return audit, nil
}

// UpdateWarrantyAuditService updates an existing warranty audit in the database.
func UpdateWarrantyAuditService(ctx context.Context, db *gorm.DB, audit *models.WarrantyAudit) (*models.WarrantyAudit, error) {
	if audit == nil || audit.ID == 0 {
		return nil, ErrMissingWarrantyAuditData
	}

	// Implement validation logic for warranty audit data before update
	// (e.g., check for required fields)
	// if err := ValidateWarrantyAudit(audit); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Save(&audit)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWarrantyAuditNotFound
		}
		return nil, fmt.Errorf("failed to update warranty audit: %w", err)
	}

	return audit, nil
}

// DeleteWarrantyAuditService deletes a warranty audit by ID from the database.
func DeleteWarrantyAuditService(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&models.WarrantyAudit{}, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrWarrantyAuditNotFound
		}
		return fmt.Errorf("failed to delete warranty audit: %w", err)
	}

	return nil
}
