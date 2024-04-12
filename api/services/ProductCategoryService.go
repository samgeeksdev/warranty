package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	ErrProductCategoryNotFound    = errors.New("product category not found")
	ErrMissingProductCategoryData = errors.New("missing product category data")
	// Add other potential errors related to product categories (e.g., ErrDuplicateProductCategory, ErrInvalidProductCategoryName)
)

// GetAllProductCategoriesService retrieves all product categories from the database.
func GetAllProductCategoriesService(ctx context.Context, db *gorm.DB) ([]models.ProductCategory, error) {
	var categories []models.ProductCategory
	result := db.WithContext(ctx).Find(&categories)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve product categories: %w", err)
	}

	return categories, nil
}

// GetProductCategoryByIDService retrieves a product category by its ID from the database.
func GetProductCategoryByIDService(ctx context.Context, db *gorm.DB, id uint) (*models.ProductCategory, error) {
	var category models.ProductCategory
	result := db.WithContext(ctx).Where("id = ?", id).First(&category)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductCategoryNotFound
		}
		return nil, fmt.Errorf("failed to retrieve product category: %w", err)
	}

	return &category, nil
}

// CreateProductCategoryService creates a new product category in the database.
func CreateProductCategoryService(ctx context.Context, db *gorm.DB, productCategory *models.ProductCategory) (*models.ProductCategory, error) {
	if productCategory == nil {
		return nil, ErrMissingProductCategoryData
	}

	// Implement validation logic for product category data before creation (e.g., name format)
	// if err := ValidateProductCategory(productCategory); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&productCategory)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create product category: %w", err)
	}

	return productCategory, nil
}

// UpdateProductCategoryService updates an existing product category in the database.
func UpdateProductCategoryService(ctx context.Context, db *gorm.DB, productCategory *models.ProductCategory) (*models.ProductCategory, error) {
	if productCategory == nil || productCategory.ID == 0 {
		return nil, ErrMissingProductCategoryData
	}

	// Implement validation logic for product category data before update (e.g., name format)
	// if err := ValidateProductCategory(productCategory); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Save(&productCategory)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductCategoryNotFound
		}
		return nil, fmt.Errorf("failed to update product category: %w", err)
	}

	return productCategory, nil
}

// DeleteProductCategoryService deletes a product category by its ID from the database.
func DeleteProductCategoryService(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&models.ProductCategory{}, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductCategoryNotFound
		}
		return fmt.Errorf("failed to delete product category: %w", err)
	}

	return nil
}
