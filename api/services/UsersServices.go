package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"warranty/api/errs"
	"warranty/api/models"
	"warranty/utilities"
)

func CreateUserService(ctx context.Context, db *gorm.DB, user *models.User) (*models.User, error) {
	if err := ValidateUser(user); err != nil {
		return nil, err
	}

	result := db.WithContext(ctx).Create(&user)
	if err := result.Error; err != nil {
		if IsUniqueConstraintError(err, "users", "username") || IsUniqueConstraintError(err, "users", "email") {
			return nil, errs.ErrUsernameOrEmailExists // Use custom error for clarity
		}
		return nil, fmt.Errorf("failed to create user: %w", err) // Wrap general errors
	}

	return user, nil
}

func GetUserByIDService(ctx context.Context, userID uint, db *gorm.DB) (*models.User, error) {
	var user models.User
	result := db.WithContext(ctx).Where("id = ?", userID).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		// Log the error for further investigation
		log.Printf("Error fetching user (ID: %d): %v\n", userID, result.Error)
		return nil, errs.ErrUserNotFound // Wrap in custom error for handling
	}

	return &user, nil
}

func UpdateUserService(ctx context.Context, db *gorm.DB, user *models.User) (*models.User, error) {
	if err := ValidateUser(user); err != nil {
		return nil, err
	}

	if user.PasswordHash != "" { // Only update password if provided
		hashedPassword, err := utilities.HashPassword(user.PasswordHash)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = hashedPassword
	}

	result := db.WithContext(ctx).Save(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func DeleteUserService(db *gorm.DB, userID uint) error {
	result := db.Delete(&models.User{}, userID)
	if err := result.Error; err != nil {
		return err // Propagate the error for handling in the controller layer
	}

	return nil // User deleted successfully
}

// unsiful
func LoginByEmailService(ctx context.Context, db *gorm.DB, email string, password string) (*models.User, error) {
	var user models.User
	result := db.WithContext(ctx).Where("email = ?", email).First(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // Handle record not found error
			return nil, errors.New(errs.UserErrorMessages["InavlidEmailOrPass"])
		}
		return nil, err
	}

	if !utilities.ComparePassword(user.PasswordHash, password) {
		return nil, errors.New(errs.UserErrorMessages["InavlidEmailOrPass"])
	}

	return &user, nil
}

func LoginByPhoneService(ctx context.Context, db *gorm.DB, phone string, verificationCode string) (*models.User, error) {
	// Implement verification code validation logic here (e.g., check against stored code)
	// ...

	var user models.User
	result := db.WithContext(ctx).Where("phone = ?", phone).First(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(errs.UserErrorMessages["InvalidPhone"])
		}
		return nil, err
	}

	return &user, nil
}

func ListUsersService(ctx context.Context, db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.WithContext(ctx).Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func LoginByUsernameService(ctx context.Context, db *gorm.DB, username string, password string) (*models.User, error) {
	var user models.User
	result := db.WithContext(ctx).Where("username = ?", username).First(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // Handle record not found error
			return nil, errors.New(errs.UserErrorMessages["InvalidUsernameOrPassword"])
		}
		return nil, err
	}

	if !utilities.ComparePassword(user.PasswordHash, password) {
		return nil, errors.New(errs.UserErrorMessages["InvalidUsernameOrPassword"])
	}

	return &user, nil
}
func FindUserByIDService(ctx context.Context, db *gorm.DB, userID uint) (*models.User, error) {
	var user models.User
	// Find the user record in the database by ID
	result := db.WithContext(ctx).First(&user, userID)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func FindUserByUsernameService(ctx context.Context, db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	// Find the user record in the database by username
	result := db.WithContext(ctx).Where("username = ?", username).First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}
