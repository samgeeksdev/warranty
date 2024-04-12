package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	ErrProductCategoryAssociationNotFound    = errors.New("product category association not found")
	ErrMissingProductCategoryAssociationData = errors.New("missing product category association data")
	// Add other potential errors related to product category associations (e.g., ErrDuplicateProductCategoryAssociation)
)

// GetAllProductCategoryAssociationsService retrieves all product category associations from the database.
func GetAllProductCategoryAssociationsService(ctx context.Context, db *gorm.DB) ([]models.ProductCategoryAssociation, error) {
	var associations []models.ProductCategoryAssociation
	result := db.WithContext(ctx).Find(&associations)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve product category associations: %w", err)
	}

	return associations, nil
}

// GetProductCategoryAssociationService retrieves a product category association by product ID and category ID from the database.
func GetProductCategoryAssociationService(ctx context.Context, db *gorm.DB, productID, categoryID uint) (*models.ProductCategoryAssociation, error) {
	var association models.ProductCategoryAssociation
	result := db.WithContext(ctx).Where("product_id = ? AND category_id = ?", productID, categoryID).First(&association)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductCategoryAssociationNotFound
		}
		return nil, fmt.Errorf("failed to retrieve product category association: %w", err)
	}

	return &association, nil
}

// CreateProductCategoryAssociationService creates a new product category association in the database.
func CreateProductCategoryAssociationService(ctx context.Context, db *gorm.DB, association *models.ProductCategoryAssociation) (*models.ProductCategoryAssociation, error) {
	if association == nil {
		return nil, ErrMissingProductCategoryAssociationData
	}

	// Implement validation logic for association data before creation (e.g., product and category existence)
	// if err := ValidateProductCategoryAssociation(association); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&association)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create product category association: %w", err)
	}

	return association, nil
}

// DeleteProductCategoryAssociationService deletes a product category association by product ID and category ID from the database.
func DeleteProductCategoryAssociationService(ctx context.Context, db *gorm.DB, productID, categoryID uint) error {
	result := db.WithContext(ctx).Delete(&models.ProductCategoryAssociation{}, map[string]interface{}{"product_id": productID, "category_id": categoryID})
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductCategoryAssociationNotFound
		}
		return fmt.Errorf("failed to delete product category association: %w", err)
	}

	return nil
}
