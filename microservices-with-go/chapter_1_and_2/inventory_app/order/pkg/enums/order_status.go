package enums

type OrderStatus int

const (
	OrderStatusPending   = OrderStatus(iota)
	OrderStatusCompleted = OrderStatus(iota)
	OrderStatusCancelled = OrderStatus(iota)
)
