package routes

import (
	"github.com/gin-gonic/gin"
	"warranty/api/handlers"
	"warranty/api/middleware" // Assuming middleware package is defined
)

func UserRoutes(router *gin.Engine) { // Receive router instance as parameter
	users := router.Group("/users")
	{
		// Registration
		users.POST("/register", handlers.RegisterUserHandler)

		// Authentication
		users.POST("/login", handlers.LoginHandlerHandler)

		// Protected routes (require authentication middleware)
		protected := users.Group("/", middleware.AuthorizeJWT()) // Ensure middleware is defined
		{
			protected.GET("/:id", handlers.GetUserByIDHandler)

			protected.GET("/me", handlers.GetLoggedInUserHandler)
			protected.PUT("/:id", handlers.UpdateUserHandler)
			protected.DELETE("/:id", handlers.DeleteUserHandler)

			// Additional routes for managing user data
			//			protected.PUT("/:id/password", handlers.ChangePassword)
			// ... other routes as needed
		}
	}
}
