package requests

type OrderItemRequest struct {
	FoodID   string `json:"foodID" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}
