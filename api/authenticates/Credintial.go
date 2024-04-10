package authenticates

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"os"
	"time"
	"warranty/api/errs"
	"warranty/api/models"
	"warranty/database"
	"warranty/helpers"
	"warranty/utilities"
)

func ValidateLoginCredentials(credentials models.LoginCredentials) error {
	// Implement validation logic for username and password (e.g., length checks, format checks)
	if len(credentials.Username) < 3 || len(credentials.Password) < 8 {
		return fmt.Errorf(errs.UserErrorMessages["UsernamePasswordShort"])
	}
	return nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
	// Fetch user from the database using your database library (e.g., GORM)
	// Example using GORM:
	db := database.GetDB() // Assuming a function to get a database connection
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf(errs.UserErrorMessages["404NotFound"])
	}
	fmt.Println(username, password, user.PasswordHash)

	// Compare the hashed password with the provided password
	if err := utilities.CheckPasswordHash(password, user.PasswordHash); err != nil {
		// Handle password mismatch or other errors
		return nil, fmt.Errorf("%s: %w", errs.UserErrorMessages["InvalidUsernameOrPassword"], err)

	}

	return &user, nil
}

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	helpers.LoadGodotEnv()
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("missing required environment variable 'SECRET_KEY'")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		// Wrap the error with more context
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}
