package database

import (
	"database/sql"
	"fmt"
	"warranty/helpers"
)

var migrationOrder = []string{
	"database/tables/manufacturers.sql",
	"database/tables/warranty_types.sql",
	"database/tables/permissions.sql",
	"database/tables/roles.sql",
	"database/tables/users.sql",
	"database/tables/user_roles.sql",
	"database/tables/product_categories.sql",
	"database/tables/products.sql",
	"database/tables/product_category_associations.sql",
	"database/tables/customers.sql",
	"database/tables/warranties.sql",
	"database/tables/warranty_audits.sql",
	"database/tables/claim_statuses.sql",
	"database/tables/claims.sql",
	"database/tables/claim_audits.sql",
	"database/tables/repair_representatives.sql",
	"database/tables/availabilities.sql",
	"database/tables/skills.sql",
	"database/tables/repair_representative_skills.sql",
}

func MigrateAll() {
	// Connect to the database
	db, err := Connect()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return // Exit program if connection fails
	}

	// Defer database closing to ensure it happens even on errors
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	// Migrate tables in the specified order
	for _, filePath := range migrationOrder {
		err := migrateTable(db, filePath)
		if err != nil {
			fmt.Println(filePath, "Error migrating table:", err)
			fmt.Println(filePath, "**Detailed Error:**", err.Error())
			return // Abort program execution on any error
		}
		fmt.Println("Migration for", filePath, "successful!")
	}
}

// migrateTable migrates a single table using the provided SQL file path
func migrateTable(db *sql.DB, filePath string) error {
	return helpers.MigrateFromFile(db, filePath)
}
