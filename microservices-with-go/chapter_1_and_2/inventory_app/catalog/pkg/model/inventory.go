package model

// CategoryID defines the unique identifier for a category.
type CategoryID int

// Category represents a product category.
type Category struct {
	ID   CategoryID `json:"id"`
	Name string     `json:"name"`
}

// SubCategoryID defines the unique identifier for a sub-category.
type SubCategoryID int

// SubCategoryBaseInfo contains the common fields for sub-category structures.
type SubCategoryBaseInfo struct {
	ID   SubCategoryID `json:"id"`
	Name string        `json:"name"`
}

// SubCategoryBasic represents the basic sub-category information used for
// create/update operations and internal DB mappings.
type SubCategoryBasic struct {
	SubCategoryBaseInfo
	CatID CategoryID `json:"categoryID"` // Renamed json tag for consistency
}

// SubCategoryDetails represents the sub-category information along with its
// associated category details. Used in GET/read operations.
type SubCategoryDetails struct {
	SubCategoryBaseInfo
	Category *Category `json:"category"` // Renamed json tag to be more intuitive
}

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
