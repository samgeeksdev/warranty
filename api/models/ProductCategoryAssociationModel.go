package models

type ProductCategoryAssociation struct {
	ProductID  int64 `json:"product_id"`
	CategoryID int64 `json:"category_id"`
	IsPrimary  bool  `json:"is_primary"`
}
