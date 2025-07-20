package enums

type OrderType int

const (
	OrderTypeSale   = OrderType(iota) // Represents a sale transaction
	OrderTypeBuy    = OrderType(iota) // Represents a purchase transaction
	OrderTypeReturn = OrderType(iota) // Represents a return transaction
)
