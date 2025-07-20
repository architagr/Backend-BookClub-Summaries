package ginhandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	catalogModel "inventory.com/catalog/pkg/model"
	"inventory.com/order/pkg/enums"
	"inventory.com/order/pkg/model"
)

// IOrderController defines the interface for order operations.
type IOrderController interface {
	CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error)
	GetAllOrders(ctx context.Context) ([]*model.Order, error)
	GetOrdersByProductID(ctx context.Context, productID catalogModel.ProductID) ([]*model.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID model.OrderID, status enums.OrderStatus) error
	GetOrder(ctx context.Context, orderID model.OrderID) (*model.Order, error)
	CurrentStock(ctx context.Context, productID catalogModel.ProductID) (int, error)
}

type orderHandler struct {
	ctrl IOrderController
}

// NewOrderHandler creates a new OrderHandler with the provided controller.
func newOrderHandler(ctrl IOrderController) *orderHandler {
	return &orderHandler{
		ctrl: ctrl,
	}
}

// CreateOrder handles the creation of a new order.
// It accepts an Order model, validates it, and then calls the controller to save it.
func (h *orderHandler) CreateOrder(ctx *gin.Context) {
	order := &model.Order{}
	if err := ctx.ShouldBindJSON(order); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid order data"})
		return
	}

	createdOrder, err := h.ctrl.CreateOrder(ctx.Request.Context(), order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdOrder)
}

// GetAllOrders retrieves all orders from the controller.
func (h *orderHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := h.ctrl.GetAllOrders(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(orders) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No orders found"})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (h *orderHandler) GetOrdersByProductID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	orders, err := h.ctrl.GetOrdersByProductID(ctx.Request.Context(), catalogModel.ProductID(productID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(orders) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No orders found for this product"})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (h *orderHandler) UpdateOrderStatusCompleted(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.ctrl.UpdateOrderStatus(ctx.Request.Context(), model.OrderID(orderID), enums.OrderStatusCompleted)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *orderHandler) UpdateOrderStatusCancelled(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.ctrl.UpdateOrderStatus(ctx.Request.Context(), model.OrderID(orderID), enums.OrderStatusCancelled)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *orderHandler) GetOrder(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.ctrl.GetOrder(ctx.Request.Context(), model.OrderID(orderID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if order == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}
	ctx.JSON(http.StatusOK, order)
}
func (h *orderHandler) CurrentStock(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	stock, err := h.ctrl.CurrentStock(ctx.Request.Context(), catalogModel.ProductID(productID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"currentStock": stock})
}
func RegisterOrderRoutes(router *gin.Engine, ctrl IOrderController) {
	handler := newOrderHandler(ctrl)

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("/", handler.CreateOrder)
		orderGroup.GET("/", handler.GetAllOrders)
		orderGroup.GET("/product/:productID", handler.GetOrdersByProductID)
		orderGroup.PUT("/:orderID/status/completed", handler.UpdateOrderStatusCompleted)
		orderGroup.PUT("/:orderID/status/cancelled", handler.UpdateOrderStatusCancelled)
		orderGroup.GET("/:orderID", handler.GetOrder)
		orderGroup.GET("/product/:productID/stock", handler.CurrentStock)
	}
}
