package model

// CategoryID defines the unique identifier for a category.
type CategoryID int

// Category represents a product category.
type Category struct {
	ID   CategoryID `json:"id"`
	Name string     `json:"name"`
}
