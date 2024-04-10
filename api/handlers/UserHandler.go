package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"warranty/api/authenticates"
	"warranty/api/controllers"
	"warranty/api/errs"
	"warranty/api/models"
	"warranty/api/services"
	"warranty/database"
	"warranty/utilities"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.UserErrorMessages["InvalidUserData"]})
		return
	}

	hashedPassword, err := utilities.HashPassword(user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.UserErrorMessages["FailedToHashing"]})
		return
	}
	user.PasswordHash = hashedPassword

	newUser, err := services.CreateUserService(ctx, database.GetDB(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func LoginHandlerHandler(c *gin.Context) {
	var credentials models.LoginCredentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.UserErrorMessages["InvalidLoginCredentials"]})
		return
	}

	// Validate username and password
	if err := authenticates.ValidateLoginCredentials(credentials); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user (replace with your authentication logic)
	user, err := authenticates.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errs.UserErrorMessages["InvalidUsernameOrPassword"]})
		return
	}

	// Generate a token (replace with your token generation logic)
	tokenString, err := authenticates.GenerateToken(uint(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.UserErrorMessages["FailedToGenerateToken"]})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetUserByIDHandler(c *gin.Context) {
	uc := controllers.UserController{}
	uc.GetUserByIDController(c)
}
func GetLoggedInUserHandler(c *gin.Context) {
	// Access user ID from context set by the middleware
	userID := c.MustGet("user_id").(int64)

	// Fetch user details using the service layer, providing a context for potential cancellation
	user, err := services.GetUserByIDService(c.Request.Context(), uint(userID), database.GetDB()) // Convert int64 to string for GetUserById
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": errs.UserErrorMessages["UserNotFound"]})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			// Log the error for debugging
			// log.Printf("Error retrieving user: %v", err)
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context) {
	uc := controllers.UserController{}
	uc.UpdateUserController(c)
}

// DeleteUser (fixed and using UserID)

// DeleteUser (fixed and using UserID)
func DeleteUserHandler(c *gin.Context) {
	// Access user ID from context set by the middleware
	userID := c.MustGet("user_id").(int64)

	// Get the user trying to delete the account
	currentUser, err := services.GetUserByIDService(c.Request.Context(), uint(userID), database.GetDB())
	if err != nil {
		// Handle errors with proper status codes and messages
		if errors.Is(err, errs.UnAuthorize) {
			c.JSON(http.StatusNotFound, gin.H{"error": errs.UserErrorMessages["UserNotFound"]})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			log.Printf("Error retrieving user: %v\n", err) // Log the error for debugging

		}
		return
	}

	// Get the user ID from the path parameter (assuming ID is in the path)
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64) // Convert string ID to int64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the user is trying to delete themselves
	if int64(currentUser.ID) != idInt {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own account"})
		return
	}

	// Call the service layer to delete the user
	err = services.DeleteUserService(database.GetDB(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusNoContent, nil) //  No content to return on successful deletion
}
