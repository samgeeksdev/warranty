package routes

import (
	"github.com/gin-gonic/gin"
	"warranty/api/handlers"
)

func WarrantyRoutes(router *gin.Engine) { // Receive router instance as parameter
	warranties := router.Group("/warranties")

	// Registration

	warranties.GET("/", handlers.GetAllWarranties)
	warranties.GET("/:id", handlers.GetWarrantyByID)
	warranties.POST("/", handlers.CreateWarranty)
	warranties.PUT("/:id", handlers.UpdateWarranty)
	warranties.DELETE("/:id", handlers.DeleteWarranty)

}
