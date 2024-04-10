package models

//
//type Manufacturer struct {
//	ID          int64  `json:"id"`
//	Name        string `json:"name"`
//	Description string `json:"description"`
//}
//
//type WarrantyType struct {
//	ID   int64  `json:"id"`
//	Name string `json:"name"`
//}
//
//type Permission struct {
//	ID          int64     `json:"id"`
//	Name        string    `json:"name"`
//	Description string    `json:"description"`
//	CreatedAt   time.Time `json:"created_at"`
//}
//
//type Role struct {
//	ID          int64     `json:"id"`
//	Name        string    `json:"name"`
//	Description string    `json:"description"`
//	CreatedAt   time.Time `json:"created_at"`
//}
//
//type User struct {
//	ID           int64      `json:"id"`
//	Username     string     `json:"username"`
//	Email        string     `json:"email"`
//	PasswordHash string     `json:"-"`
//	FirstName    string     `json:"first_name"`
//	LastName     string     `json:"last_name"`
//	PhoneNumber  string     `json:"phone_number"`
//	CreatedAt    time.Time  `json:"created_at"`
//	UpdatedAt    *time.Time `json:"updated_at"`
//	Active       bool       `json:"active"`
//}
//
//type UserRole struct {
//	ID        int64     `json:"id"`
//	UserID    int64     `json:"user_id"`
//	RoleID    int       `json:"role_id"`
//	CreatedAt time.Time `json:"created_at"`
//}
//
//type ProductCategory struct {
//	ID          int64      `json:"id"`
//	Name        string     `json:"name"`
//	Description string     `json:"description"`
//	CreatedAt   time.Time  `json:"created_at"`
//	UpdatedAt   *time.Time `json:"updated_at"`
//}
//
//type Product struct {
//	ID             int64      `json:"id"`
//	Name           string     `json:"name"`
//	Slug           string     `json:"slug"`
//	Description    string     `json:"description"`
//	CategoryID     int        `json:"category_id"`
//	ManufacturerID *int       `json:"manufacturer_id"`
//	Model          string     `json:"model"`
//	SKU            string     `json:"sku"`
//	Price          float64    `json:"price"`
//	Stock          int        `json:"stock"`
//	CreatedAt      time.Time  `json:"created_at"`
//	UpdatedAt      *time.Time `json:"updated_at"`
//	WarrantyMonths *int       `json:"warranty_months"`
//	Retired        bool       `json:"retired"`
//}
//
//type ProductCategoryAssociation struct {
//	ProductID  int64 `json:"product_id"`
//	CategoryID int64 `json:"category_id"`
//	IsPrimary  bool  `json:"is_primary"`
//}
//
//type Customer struct {
//	ID          int64     `json:"id"`
//	UserID      *int64    `json:"user_id"`
//	Name        string    `json:"name"`
//	Email       string    `json:"email"`
//	PhoneNumber string    `json:"phone_number"`
//	Address     string    `json:"address"`
//	CreatedAt   time.Time `json:"created_at"`
//}
//
//type Warranty struct {
//	ID             int64      `json:"id"`
//	ProductID      int64      `json:"product_id"`
//	WarrantyTypeID int        `json:"warranty_type_id"`
//	DurationMonths int        `json:"duration_months"`
//	StartDate      time.Time  `json:"start_date"`
//	EndDate        *time.Time `json:"end_date"`
//	RegisteredBy   int64      `json:"registered_by"`
//	CustomerID     *int64     `json:"customer_id"`
//}
//
//type WarrantyAudit struct {
//	ID            int64     `json:"id"`
//	WarrantyID    int64     `json:"warranty_id"`
//	ModifiedField string    `json:"modified_field"`
//	OldValue      string    `json:"old_value"`
//	NewValue      string    `json:"new_value"`
//	ModifiedBy    int64     `json:"modified_by"`
//	ModifiedAt    time.Time `json:"modified_at"`
//}
//
//type ClaimStatus struct {
//	ID   int64  `json:"id"`
//	Name string `json:"name"`
//}
//
//type Claim struct {
//	ID          int64     `json:"id"`
//	WarrantyID  int64     `json:"warranty_id"`
//	Description string    `json:"description"`
//	ClaimDate   time.Time `json:"claim_date"`
//	StatusID    int       `json:"status_id"`
//	Resolution  string    `json:"resolution"`
//	ClaimFiles  string    `json:"claim_files_json"`
//}
//
//type ClaimAudit struct {
//	ID            int64     `json:"id"`
//	ClaimID       int64     `json:"claim_id"`
//	ModifiedField string    `json:"modified_field"`
//	OldValue      string    `json:"old_value"`
//	NewValue      string    `json:"new_value"`
//	ModifiedBy    int64     `json:"modified_by"`
//	ModifiedAt    time.Time `json:"modified_at"`
//}
//
//type RepairRepresentative struct {
//	ID             int64      `json:"id"`
//	UserID         *int64     `json:"user_id"`
//	Name           string     `json:"name"`
//	PhoneNumber    string     `json:"phone_number"`
//	Email          string     `json:"email"`
//	Active         bool       `json:"active"`
//	Certification  string     `json:"certification"`
//	Specialization string     `json:"specialization"`
//	CreatedAt      time.Time  `json:"created_at"`
//	UpdatedAt      *time.Time `json:"updated_at"`
//}
//
//type Availability struct {
//	ID                     int64      `json:"id"`
//	RepairRepresentativeID int64      `json:"repair_representative_id"`
//	AvailabilityType       string     `json:"availability_type"`
//	AvailabilityDetails    string     `json:"availability_details"`
//	CreatedAt              time.Time  `json:"created_at"`
//	UpdatedAt              *time.Time `json:"updated_at"`
//}
//
//type Skill struct {
//	ID          int64      `json:"id"`
//	Name        string     `json:"name"`
//	Description string     `json:"description"`
//	CreatedAt   time.Time  `json:"created_at"`
//	UpdatedAt   *time.Time `json:"updated_at"`
//}
//
//type RepairRepresentativeSkill struct {
//	RepairRepresentativeID int64 `json:"repair_representative_id"`
//	SkillID                int64 `json:"skill_id"`
//}
