package routes

import (
	"github.com/gin-gonic/gin"
	"warranty/api/handlers"
)

func Routes() {
	router := gin.Default()

	// User routes
	users := router.Group("/users")
	{
		users.POST("/register", handlers.RegisterUserHandler)
		users.POST("/login", handlers.LoginHandlerHandler)
		users.GET("/:id", handlers.GetUserByIDHandler)
		users.PUT("/:id", handlers.UpdateUserHandler)
		users.DELETE("/:id", handlers.DeleteUserHandler)
	}

	// Product routes
	products := router.Group("/products")
	{
		products.GET("/", handlers.GetProducts)
		products.POST("/", handlers.CreateProduct)
		products.GET("/:id", handlers.GetProductByID)
		products.PUT("/:id", handlers.UpdateProduct)
		products.DELETE("/:id", handlers.DeleteProduct)
	}

	// User routes

	// Role routes
	roles := router.Group("/roles")
	{
		roles.GET("/", handlers.GetRoles)
		roles.POST("/", handlers.CreateRole)
		roles.GET("/:id", handlers.GetRoleByID)
		roles.PUT("/:id", handlers.UpdateRole)
		roles.DELETE("/:id", handlers.DeleteRole)
	}

	// User role routes
	userRoleRoutes := router.Group("/user_roles")
	{
		userRoleRoutes.GET("/", handlers.GetAllUserRoles)
		userRoleRoutes.GET("/:user_id/:role_id", handlers.GetUserRole)
		userRoleRoutes.POST("/", handlers.CreateUserRole)
		userRoleRoutes.DELETE("/:user_id/:role_id", handlers.DeleteUserRole)
	}

	// Product category routes
	productCategories := router.Group("/product-categories")
	{
		productCategories.GET("/", handlers.GetProductCategories)
		productCategories.POST("/", handlers.CreateProductCategory)
		productCategories.GET("/:id", handlers.GetProductCategoryByID)
		productCategories.PUT("/:id", handlers.UpdateProductCategory)
		productCategories.DELETE("/:id", handlers.DeleteProductCategory)
	}

	// Product category association routes
	categoryAssociationRoutes := router.Group("/product_category_associations")
	{
		categoryAssociationRoutes.GET("/", handlers.GetAllProductCategoryAssociations)
		categoryAssociationRoutes.GET("/:product_id/:category_id", handlers.GetProductCategoryAssociation)
		categoryAssociationRoutes.POST("/", handlers.CreateProductCategoryAssociation)
		categoryAssociationRoutes.DELETE("/:product_id/:category_id", handlers.DeleteProductCategoryAssociation)
	}

	// Warranty routes
	warrantyRoutes := router.Group("/warranties")
	{
		warrantyRoutes.GET("/", handlers.GetAllWarranties)
		warrantyRoutes.GET("/:id", handlers.GetWarrantyByID)
		warrantyRoutes.POST("/", handlers.CreateWarranty)
		warrantyRoutes.PUT("/:id", handlers.UpdateWarranty)
		warrantyRoutes.DELETE("/:id", handlers.DeleteWarranty)
	}

	// Warranty audit routes
	warrantyAuditRoutes := router.Group("/warranty_audits")
	{
		warrantyAuditRoutes.GET("/", handlers.GetAllWarrantyAudits)
		warrantyAuditRoutes.GET("/:id", handlers.GetWarrantyAuditByID)
		warrantyAuditRoutes.POST("/", handlers.CreateWarrantyAudit)
		warrantyAuditRoutes.PUT("/:id", handlers.UpdateWarrantyAudit)
		warrantyAuditRoutes.DELETE("/:id", handlers.DeleteWarrantyAudit)
	}

	// Claim routes
	claimRoutes := router.Group("/claims")
	{
		claimRoutes.GET("/", handlers.GetAllClaims)
		claimRoutes.GET("/:id", handlers.GetClaimByID)
		claimRoutes.POST("/", handlers.CreateClaim)
		claimRoutes.PUT("/:id", handlers.UpdateClaim)
		claimRoutes.DELETE("/:id", handlers.DeleteClaim)
	}

	// Claim audit routes
	claimAuditRoutes := router.Group("/claim_audits")
	{
		claimAuditRoutes.GET("/", handlers.GetAllClaimAudits)
		claimAuditRoutes.GET("/:id", handlers.GetClaimAuditByID)
		claimAuditRoutes.POST("/", handlers.CreateClaimAudit)
		claimAuditRoutes.PUT("/:id", handlers.UpdateClaimAudit)
		claimAuditRoutes.DELETE("/:id", handlers.DeleteClaimAudit)
	}
	// manufacturers routes
	// Manufacturers
	manufacturers := router.Group("/manufacturers")
	{
		manufacturers.GET("/", handlers.GetManufacturers)
		manufacturers.POST("/", handlers.CreateManufacturer)
		manufacturers.GET("/:id", handlers.GetManufacturerByID)
		manufacturers.PUT("/:id", handlers.UpdateManufacturer)
		manufacturers.DELETE("/:id", handlers.DeleteManufacturer)
	}

	// Warranty Types
	warrantyTypes := router.Group("/warranty-types")
	{
		warrantyTypes.GET("/", handlers.GetWarrantyTypes)
		warrantyTypes.POST("/", handlers.CreateWarrantyType)
		warrantyTypes.GET("/:id", handlers.GetWarrantyTypeByID)
		warrantyTypes.PUT("/:id", handlers.UpdateWarrantyType)
		warrantyTypes.DELETE("/:id", handlers.DeleteWarrantyType)
	}

	// Permissions
	permissions := router.Group("/permissions")
	{
		permissions.GET("/", handlers.GetPermissions)
		permissions.POST("/", handlers.CreatePermission)
		permissions.GET("/:id", handlers.GetPermissionByID)
		permissions.PUT("/:id", handlers.UpdatePermission)
		permissions.DELETE("/:id", handlers.DeletePermission)
	}

	// Roles

	// Users

	// User Roles
	userRoles := router.Group("/users/:user_id/roles")
	{
		userRoles.GET("/", handlers.GetUserRoles)
		userRoles.POST("/", handlers.AssignUserRole)
		userRoles.DELETE("/:role_id", handlers.RemoveUserRole)
	}

	// Product Categories

	// Products

	// Customers
	customers := router.Group("/customers")
	{
		customers.GET("/", handlers.GetCustomers)
		customers.POST("/", handlers.CreateCustomer)
		customers.GET("/:id", handlers.GetCustomerByID)
		customers.PUT("/:id", handlers.UpdateCustomer)
		customers.DELETE("/:id", handlers.DeleteCustomer)
	}

	// Run the server
	router.Run(":8080")
}
