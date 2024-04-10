// controllers/user_controller.go

package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"warranty/api/errs"
	"warranty/api/models"
	"warranty/api/services"
	"warranty/database"
	"warranty/utilities"
)

type UserController struct{}

func (uc *UserController) RegisterUserController(c *gin.Context) {
	ctx := c.Request.Context()

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.UserErrorMessages["InvalidUserData"]})
		return
	}

	// Hash the password
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

func (uc *UserController) GetUserByIDController(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("id")

	userIDInt, err := strconv.Atoi(userId)
	if err != nil {
		// Handle conversion error (e.g., invalid format)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := services.GetUserByIDService(ctx, uint(userIDInt), database.GetDB())
	if err != nil {
		if err == errs.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": errs.UserErrorMessages["UserNotFound"]})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUserController(c *gin.Context) {
	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.UserErrorMessages["InvalidUserData"]})
		return
	}

	userId := c.Param("id")
	// Convert string ID to int64
	var err error
	updatedUser.ID, err = strconv.ParseInt(userId, 10, 64) // Assuming ID is base 10

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	ctx := c.Request.Context()
	db := database.GetDB()

	// Call the service layer to update the user
	user, err := services.UpdateUserService(ctx, db, &updatedUser)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": errs.UserErrorMessages["UserNotFound"]})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// User updated successfully
	fmt.Println(user)
	c.JSON(http.StatusOK, user)
}
