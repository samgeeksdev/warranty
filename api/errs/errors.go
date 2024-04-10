package errs

import (
	"errors"
)

var ErrUsernameOrEmailExists = errors.New("username or email already exists")
var ErrUserNotFound = errors.New("UserNotFound")
var UnAuthorize = errors.New("PLease First Logedin")
var ErrMissingField = errors.New("ErrMissingField")

//
//const (
//	// Common error codes
//	ErrBadRequest          = 400 // Bad request format
//	ErrUnauthorized        = 401 // Invalid or missing credentials
//	ErrForbidden           = 403 // User lacks permission
//	ErrNotFound            = 404 // Resource not found
//	ErrMethodNotAllowed    = 405 // Invalid HTTP method
//	ErrInternalServerError = 500 // Internal server error
//
//	// LoginHandler specific error codes
//	ErrInvalidLoginCredentials   = 1001
//	ErrUserNotFound              = 1002
//	ErrInvalidUsernameOrPassword = 1003
//
//	// RegisterUser specific error codes
//	ErrInvalidUserData       = 1010
//	ErrPasswordHashingFailed = 1011
//	ErrDatabaseOperation     = 1012
//)
//const (
//	UsernameRequire = "username is required"
//	UsernameLen     = "username is required"
//)
//
//// Error messages (consider using a map for more flexibility if needed)
//var UserErrorMessages = map[int]string{
//
//	ErrBadRequest:          "Invalid request format",
//	ErrUnauthorized:        "Invalid or missing credentials",
//	ErrForbidden:           "User lacks permission",
//	ErrNotFound:            "Resource not found",
//	ErrMethodNotAllowed:    "Invalid HTTP method",
//	ErrInternalServerError: "Internal server error",
//
//	ErrInvalidLoginCredentials:   "Invalid email or password",
//	ErrUserNotFound:              "User not found",
//	ErrInvalidUsernameOrPassword: "Invalid username or password",
//
//	ErrInvalidUserData:       "Invalid user data provided",
//	ErrPasswordHashingFailed: "Failed to hash password",
//	ErrDatabaseOperation:     "Database operation failed",
//}

// Helper function to get error message by code (optional)
//func GetErrorMessage(code int) string {
//	message, ok := UserErrorMessages[code]
//	if !ok {
//		return fmt.Sprintf("Unknown error (code: %d)", code)
//	}
//	return message
//}

var UserErrorMessages = map[string]string{
	// Common errors (using HTTP status codes)
	"InvalidLoginCredentials":     "InvalidUsernameOrPassword",
	"InvalidUsernameOrPassword":   "InvalidUsernameOrPassword",
	"FailedToGenerateToken":       "FailedToGenerateToken",
	"InvalidUserData":             "InvalidUserData",
	"FailedToHashing":             "failed to hash password",
	"UsernameLengthOutOfRange":    "username must be between 3 and 20 characters",
	"UsernameContainsSpecialChar": "username cannot contain special characters or spaces",
	"InvalidEmailFormat":          "InvalidEmailFormat",
	"PasswordIsRequired":          "PasswordIsRequired",
	"PasswordTooShort":            "password must be at least 8 characters long",
	"404NotFound":                 "Username Not Found",
	"UsernamePasswordShort":       "UsernamePasswordShort",
	"UserNameRequired":            "UserName is Required",
	"UserCharSpace":               "username cannot contain special characters or spaces",
	"RequireEmail":                "email is required",
	"RequirePassword":             "password is required",
	"PasswordLen":                 "password must be at least 8 characters long",
	"ValidationFailed":            "validation failed",
	"DuplicateKey":                "duplicate key value violates unique constraint",
	"InvalidPhone":                "invalid phone number",
	"InavlidEmailOrPass":          "invalid email or password",
	"InavlidEmailFormat":          "InavlidEmailFormat",
	"UserNotFound":                "Username Not Found",

	// ... (Add other validation errors with appropriate keys and messages)
}
