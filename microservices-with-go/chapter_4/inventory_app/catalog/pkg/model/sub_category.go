package model

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
	BaseInfo SubCategoryBaseInfo `json:"baseInfo"`
	CatID    CategoryID          `json:"categoryID"` // Renamed json tag for consistency
}

// SubCategoryDetails represents the sub-category information along with its
// associated category details. Used in GET/read operations.
type SubCategoryDetails struct {
	SubCategoryBaseInfo
	Category *Category `json:"category"` // Renamed json tag to be more intuitive
}
