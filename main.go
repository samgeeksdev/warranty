package main

import (
	"github.com/gin-gonic/gin"
	"warranty/api/routes"
)

func main() {
	// Establish database connection (replace with your connection details)
	//db, err := database.GormConnect()
	//if err != nil {
	//	panic(err)
	//}

	router := gin.Default()
	routes.UserRoutes(router)
	routes.WarrantyRoutes(router)
	router.Run(":8080")

}
