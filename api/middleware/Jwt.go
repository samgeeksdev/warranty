package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"warranty/helpers"

	"github.com/gin-gonic/gin"
)

// Use a secure way to store and access the secretKey

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Check for missing token
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// Parse and validate JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			helpers.LoadGodotEnv()
			secretKey := os.Getenv("SECRET_KEY")
			if secretKey == "" {
				fmt.Println(errors.New("missing required environment variable 'SECRET_KEY'"))

				return "", errors.New("missing required environment variable 'SECRET_KEY'")
			}

			return []byte(secretKey), nil // Return byte slice for signing method
		})

		// Handle potential errors
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error parsing token"})
			return
		}

		// Validate claims
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := int64(claims["user_id"].(float64)) // Adjust claim key if needed
		c.Set("user_id", userID)
		c.Next() // Continue processing the request
	}
}
