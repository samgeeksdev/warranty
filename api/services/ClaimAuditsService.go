package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	ErrClaimAuditNotFound = errors.New("claim audit not found")
	ErrMissingAuditData   = errors.New("missing claim audit data")
)

// Services for claim audits

func GetAllClaimAuditsService(ctx context.Context, db *gorm.DB) ([]models.ClaimAudit, error) {
	var audits []models.ClaimAudit
	result := db.WithContext(ctx).Find(&audits)
	if err := result.Error; err != nil {
		return nil, err
	}

	return audits, nil
}

func GetClaimAuditByIDService(ctx context.Context, db *gorm.DB, id uint) (*models.ClaimAudit, error) {
	var audit models.ClaimAudit
	result := db.WithContext(ctx).Where("id = ?", id).First(&audit)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClaimAuditNotFound
		}
		return nil, err
	}

	return &audit, nil
}

func CreateClaimAuditService(ctx context.Context, db *gorm.DB, audit *models.ClaimAudit) (*models.ClaimAudit, error) {
	if audit == nil {
		return nil, ErrMissingAuditData
	}

	// Validate claim audit data before creation (implement validation logic)
	// if err := ValidateClaimAudit(audit); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&audit)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create claim audit: %w", err)
	}

	return audit, nil
}

func UpdateClaimAuditService(ctx context.Context, db *gorm.DB, audit *models.ClaimAudit) (*models.ClaimAudit, error) {
	if audit == nil || audit.ID == 0 {
		return nil, ErrMissingAuditData
	}

	// Validate claim audit data before update (implement validation logic)
	// if err := ValidateClaimAudit(audit); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Save(&audit)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClaimAuditNotFound
		}
		return nil, fmt.Errorf("failed to update claim audit: %w", err)
	}

	return audit, nil
}

func DeleteClaimAuditService(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&models.ClaimAudit{}, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrClaimAuditNotFound
		}
		return fmt.Errorf("failed to delete claim audit: %w", err)
	}

	return nil
}
