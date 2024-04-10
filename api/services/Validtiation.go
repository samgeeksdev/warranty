package services

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"regexp"
	"strings"
	"warranty/api/errs"
	"warranty/api/models"
)

func ValidateUser(user *models.User) error {
	var errMsgs []string // Collect error messages

	// Username validation
	if user.Username == "" {
		errMsgs = append(errMsgs, errs.UserErrorMessages["UserNameRequired"])
	} else if len(user.Username) < 3 || len(user.Username) > 20 {
		errMsgs = append(errMsgs, errs.UserErrorMessages["UsernameLengthOutOfRange"])
	} else if !IsValidUsername(user.Username) {
		errMsgs = append(errMsgs, errs.UserErrorMessages["UserCharSpace"])
	}

	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + `'{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	if user.Email == "" {
		errMsgs = append(errMsgs, "email is required")
	} else if !regexp.MustCompile(emailRegex).MatchString(user.Email) {
		errMsgs = append(errMsgs, errs.UserErrorMessages["RequireEmail"])
	}

	// Password validation (consider minimum length and complexity)
	if user.PasswordHash == "" {
		errMsgs = append(errMsgs, errs.UserErrorMessages["PasswordIsRequired"])
	} else if len(user.PasswordHash) < 8 {
		errMsgs = append(errMsgs, "password must be at least 8 characters long")
	}

	// Additional checks (optional)
	// - Check if username already exists in the database
	// - ...

	if len(errMsgs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, ", "))
	}

	return nil
}

func IsUniqueConstraintError(err error, table, column string) bool {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) && pgError.Code == "23505" && strings.Contains(pgError.Message, fmt.Sprintf(errs.UserErrorMessages["DuplicateKey"], " \"%s_%s_key\"", table, column)) {
		return true
	}
	return false
}

// isValidUsername checks if the username contains only allowed characters
func IsValidUsername(username string) bool {
	allowedChars := `[a-zA-Z0-9_.-]` // Adjust allowed characters as needed
	return regexp.MustCompile(allowedChars).MatchString(username)
}
