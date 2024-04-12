package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	// Define potential errors related to products
	ErrProductNotFound    = errors.New("product not found")
	ErrMissingProductData = errors.New("missing product data")
	// Add other errors (e.g., ErrDuplicateProduct, ErrInvalidProductName)
)

// GetProductsService retrieves a list of products from the database based on the provided filter.
func GetProductsService(ctx context.Context, db *gorm.DB, filter *models.ProductFilter) ([]models.Product, error) {
	var products []models.Product
	dbQuery := db.WithContext(ctx) // Start building the query

	if filter != nil {
		// Apply filters based on filter properties
		if filter.Name != nil && *filter.Name != "" {
			dbQuery = dbQuery.Where("name LIKE ?", "%"+*filter.Name+"%")
		}
		if filter.CategoryID != nil && *filter.CategoryID != 0 {
			dbQuery = dbQuery.Where("category_id = ?", *filter.CategoryID)
		}
		// Add logic for other filter criteria
	}

	result := dbQuery.Find(&products)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve products: %w", err)
	}

	return products, nil
}

// GetProductService retrieves a product by its ID from the database.
func GetProductService(ctx context.Context, db *gorm.DB, id uint) (*models.Product, error) {
	var product models.Product
	result := db.WithContext(ctx).Where("id = ?", id).First(&product)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to retrieve product: %w", err)
	}

	return &product, nil
}

// CreateProductService creates a new product in the database.
func CreateProductService(ctx context.Context, db *gorm.DB, product *models.Product) (*models.Product, error) {
	if product == nil {
		return nil, ErrMissingProductData
	}

	// Implement validation logic for product data before creation (e.g., name format, required fields)
	// if err := ValidateProduct(product); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&product)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

// UpdateProductService updates an existing product in the database.
func UpdateProductService(ctx context.Context, db *gorm.DB, product *models.Product) (*models.Product, error) {
	if product == nil || product.ID == 0 {
		return nil, ErrMissingProductData
	}

	// Implement validation logic for product data before update (e.g., name format)
	// if err := ValidateProduct(product); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Save(&product)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return product, nil
}

// DeleteProductService deletes a product by its ID from the database.
func DeleteProductService(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&models.Product{}, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
