package memory

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	catalogModel "inventory.com/catalog/pkg/model"
	"inventory.com/order/pkg/enums"
	"inventory.com/order/pkg/model"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type Order struct {
	mu     sync.RWMutex
	orders map[catalogModel.ProductID][]*model.Order
	seqID  int
}

func New() *Order {
	return &Order{
		orders: make(map[catalogModel.ProductID][]*model.Order),
		seqID:  0,
	}
}

// Create adds a new order to the in-memory store.
// It automatically assigns a unique ID and initializes timestamps.
// Returns the created order record.
func (repo *Order) Create(ctx context.Context, orderRecord *model.Order) (*model.Order, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.seqID++
	orderRecord.ID = model.OrderID(repo.seqID)
	if _, exists := repo.orders[orderRecord.ProductID]; !exists {
		repo.orders[orderRecord.ProductID] = []*model.Order{}
	}
	repo.orders[orderRecord.ProductID] = append(repo.orders[orderRecord.ProductID], orderRecord)
	orderRecord.CreatedAt = time.Now()
	orderRecord.UpdatedAt = orderRecord.CreatedAt

	return orderRecord, nil
}

// GetAll retrieves all orders across all products.
// Returns ErrOrderNotFound if no orders exist.
func (repo *Order) GetAll(ctx context.Context) ([]*model.Order, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var allOrders []*model.Order
	for _, orders := range repo.orders {
		allOrders = append(allOrders, orders...)
	}

	if len(allOrders) == 0 {
		return nil, ErrOrderNotFound
	}
	return allOrders, nil
}

// GetByProductID retrieves all orders for a specific product ID.
// Returns ErrOrderNotFound if no orders exist for that product.
func (repo *Order) GetByProductID(ctx context.Context, productID catalogModel.ProductID) ([]*model.Order, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	orders := repo.orders[productID]
	if len(orders) == 0 {
		return nil, fmt.Errorf("%w, productId:%d", ErrOrderNotFound, productID)
	}
	return orders, nil
}

// UpdateStatus updates the status of an order by its ID.
// Returns ErrOrderNotFound if the order does not exist.
func (repo *Order) UpdateStatus(ctx context.Context, orderID model.OrderID, status enums.OrderStatus) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, orders := range repo.orders {
		for _, order := range orders {
			if order.ID == orderID {
				order.Status = status
				// Update the timestamp to reflect the change
				order.UpdatedAt = time.Now()
				return nil
			}
		}
	}
	return fmt.Errorf("%w: id=%d", ErrOrderNotFound, orderID)
}

// Get retrieves an order by its ID. Returns ErrOrderNotFound if not found.
func (repo *Order) Get(ctx context.Context, orderID model.OrderID) (*model.Order, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, orders := range repo.orders {
		for _, order := range orders {
			if order.ID == orderID {
				return order, nil // Order found
			}
		}
	}
	return nil, fmt.Errorf("%w: id=%d", ErrOrderNotFound, orderID)
}
