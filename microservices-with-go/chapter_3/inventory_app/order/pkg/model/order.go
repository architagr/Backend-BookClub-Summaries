// Package model contains all data structures used within the Order service.
package model

import (
	"time"

	catalogModel "inventory.com/catalog/pkg/model"
	"inventory.com/order/pkg/enums"
)

// OrderID represents the unique identifier for an Order.
type OrderID int

// Order represents a customer's transaction in the system.
// It can include one or more items, and supports multiple types such as PURCHASE, SALE, or RETURN.
//
// Fields:
//   - ID: Unique identifier for the order.
//   - ProductID: Unique identifier of the product from the Catalog service.
//   - Quantity: Number of units being ordered.
//   - Price: Price per unit (or total, depending on how billing is handled).
//   - Type: Specifies the nature of the order (e.g., PURCHASE, SALE).
//   - CustomerID: Identifier of the customer placing the order (optional).
//   - CreatedAt: Timestamp of when the order was created.
//   - UpdatedAt: Timestamp of the last update made to the order.
//   - Status: Current status of the order (e.g., PENDING, COMPLETED, CANCELLED).
type Order struct {
	ID         OrderID                `json:"id"`
	ProductID  catalogModel.ProductID `json:"productID"`
	Quantity   int                    `json:"quantity"`   // Number of items
	Price      float64                `json:"price"`      // Price per unit or total, depending on use case
	Type       enums.OrderType        `json:"type"`       // e.g. PURCHASE, SALE or RETURN
	CustomerID int                    `json:"customerID"` // Optional: if you're supporting customer data
	CreatedAt  time.Time              `json:"createdAt"`  // Timestamp for auditing
	UpdatedAt  time.Time              `json:"updatedAt"`  // Useful for updates or tracking
	Status     enums.OrderStatus      `json:"status"`     // e.g. PENDING, COMPLETED, CANCELLED
}
