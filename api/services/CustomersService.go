package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"warranty/api/models"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrMissingCustomerData = errors.New("missing customer data")
	// Add other potential errors related to customers (e.g., ErrDuplicateCustomer, ErrInvalidCustomerEmail)
)

// Services for customers

func GetAllCustomersService(ctx context.Context, db *gorm.DB) ([]models.Customer, error) {
	var customers []models.Customer
	result := db.WithContext(ctx).Find(&customers)
	if err := result.Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func GetCustomerByIDService(ctx context.Context, db *gorm.DB, id uint) (*models.Customer, error) {
	var customer models.Customer
	result := db.WithContext(ctx).Where("id = ?", id).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}

	return &customer, nil
}

// CreateCustomerService creates a new customer in the database
func CreateCustomerService(ctx context.Context, db *gorm.DB, customer *models.Customer) (*models.Customer, error) {
	if customer == nil {
		return nil, ErrMissingCustomerData
	}

	// Implement validation logic for customer data before creation (e.g., email format)
	// if err := ValidateCustomer(customer); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Create(&customer)
	if err := result.Error; err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}

	return customer, nil
}

// UpdateCustomerService updates an existing customer in the database
func UpdateCustomerService(ctx context.Context, db *gorm.DB, customer *models.Customer) (*models.Customer, error) {
	if customer == nil || customer.ID == 0 {
		return nil, ErrMissingCustomerData
	}

	// Implement validation logic for customer data before update (e.g., email format)
	// if err := ValidateCustomer(customer); err != nil {
	//   return nil, err
	// }

	result := db.WithContext(ctx).Save(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, fmt.Errorf("failed to update customer: %w", err)
	}

	return customer, nil
}

// DeleteCustomerService deletes a customer from the database by ID
func DeleteCustomerService(ctx context.Context, db *gorm.DB, id uint) error {
	result := db.WithContext(ctx).Delete(&models.Customer{}, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCustomerNotFound
		}
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	return nil
}
