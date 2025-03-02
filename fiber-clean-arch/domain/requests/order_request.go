package requests

type CreateOrderRequest struct {
	OrderItems []OrderItemRequest `json:"order_items" validate:"required"`
}

type UpdateOrderStatus struct {
	OrderID string `json:"order_id" validate:"required"`
	Status  string `json:"status" validate:"required"`
}
