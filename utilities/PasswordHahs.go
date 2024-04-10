package utilities

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

var ErrPasswordMismatch = errors.New("password mismatch")

func HashPassword(password string) (string, error) {
	// Generate a salt for the hash with a cost of 10
	// Recommended cost values are between 10 and 14
	hasher := md5.New()

	// Write the password bytes to the hasher
	hasher.Write([]byte(password))

	// Get the hashed password bytes
	hashedBytes := hasher.Sum(nil)

	// Convert the hashed bytes to a hexadecimal string
	hashedPassword := hex.EncodeToString(hashedBytes)
	// Return the hashed password as a base64 encoded string
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) error {
	// Hash the provided password using MD5
	hasher := md5.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	hashed := hex.EncodeToString(hasher.Sum(nil))

	// Compare the hashed passwords
	if hashed != hashedPassword {
		return fmt.Errorf("passwords do not match")
	}

	return nil
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func HasPasswordComplexity(password string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)

	return hasUpper //&& hasLower && hasNumber //&& hasSymbol && !hasSpace
}
