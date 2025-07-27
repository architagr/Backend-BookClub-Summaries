package model

// ProductID defines the unique identifier for a product.
type ProductID int

// ProductBaseInfo contains the common fields for product structures.
type ProductBaseInfo struct {
	ID           ProductID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Manufacturer string    `json:"manufacturer"`
	ListCost     int       `json:"listCost"`
}

// ProductBasic represents the minimal product data required for
// create/update operations and DB mappings.
type ProductBasic struct {
	ProductBaseInfo
	SubCatID SubCategoryID `json:"subCategoryID"`
}

// ProductInformation represents a product along with its associated
// sub-category details. Used for read operations like GET.
type ProductInformation struct {
	ProductBaseInfo
	SubCategoryDetails *SubCategoryDetails `json:"subCategory"` // Updated to include full detail, not just ID
}
