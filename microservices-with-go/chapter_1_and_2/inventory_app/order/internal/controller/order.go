package controller

import (
	"context"
	"errors"

	catalogModel "inventory.com/catalog/pkg/model"
	"inventory.com/order/pkg/enums"
	"inventory.com/order/pkg/model"
)

type IOrderRepository interface {
	Create(ctx context.Context, orderRecord *model.Order) (*model.Order, error)
	GetAll(ctx context.Context) ([]*model.Order, error)
	GetByProductID(ctx context.Context, productID catalogModel.ProductID) ([]*model.Order, error)
	UpdateStatus(ctx context.Context, orderID model.OrderID, status enums.OrderStatus) error
	Get(ctx context.Context, orderID model.OrderID) (*model.Order, error)
}
type OrderController struct {
	repo IOrderRepository
}

// NewOrderController creates a new instance of OrderController with the provided repository.
func NewOrderController(repo IOrderRepository) *OrderController {
	return &OrderController{
		repo: repo,
	}
}

// CreateOrder handles the creation of a new order.
// It accepts an Order model, validates it, and then calls the repository to save it.
func (c *OrderController) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	if order == nil {
		return nil, errors.New("order cannot be nil")
	}
	if order.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}
	if order.Price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	// Set initial status to PENDING
	order.Status = enums.OrderStatusPending

	return c.repo.Create(ctx, order)
}

// GetAllOrders retrieves all orders from the repository.
func (c *OrderController) GetAllOrders(ctx context.Context) ([]*model.Order, error) {
	orders, err := c.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, errors.New("no orders found")
	}
	return orders, nil
}

// GetOrdersByProductID retrieves all orders for a specific product ID.
func (c *OrderController) GetOrdersByProductID(ctx context.Context, productID catalogModel.ProductID) ([]*model.Order, error) {
	if productID <= 0 {
		return nil, errors.New("invalid product ID")
	}

	orders, err := c.repo.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, errors.New("no orders found for the specified product")
	}
	return orders, nil
}

// UpdateOrderStatus updates the status of an existing order by its ID.
func (c *OrderController) UpdateOrderStatus(ctx context.Context, orderID model.OrderID, status enums.OrderStatus) error {
	if orderID <= 0 {
		return errors.New("invalid order ID")
	}
	if status < enums.OrderStatusPending || status > enums.OrderStatusCancelled {
		return errors.New("invalid order status")
	}

	return c.repo.UpdateStatus(ctx, orderID, status)
}

// GetOrder retrieves a specific order by its ID.
func (c *OrderController) GetOrder(ctx context.Context, orderID model.OrderID) (*model.Order, error) {
	if orderID <= 0 {
		return nil, errors.New("invalid order ID")
	}

	order, err := c.repo.Get(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

// CurrentStock calculates the current stock for a product based on its orders.
func (c *OrderController) CurrentStock(ctx context.Context, productID catalogModel.ProductID) (int, error) {
	orders, err := c.repo.GetByProductID(ctx, productID)
	if err != nil {
		return 0, err
	}

	totalQuantity := 0
	for _, order := range orders {
		if order.Status != enums.OrderStatusCompleted {
			continue
		}
		if order.Type == enums.OrderTypeBuy || order.Type == enums.OrderTypeReturn {
			totalQuantity += order.Quantity
		}

		if order.Type == enums.OrderTypeSale {
			totalQuantity -= order.Quantity
		}
	}

	return totalQuantity, nil
}
