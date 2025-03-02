package repositories

import (
	"context"
	"fiber-clean-arch/domain/models"
	"fiber-clean-arch/domain/requests"
)

type OrderItemRepository interface {
	GetOrderItemsByOrderID(ctx context.Context, orderID string) ([]models.OrderItem, error)
	CreateOrderItems(ctx context.Context, orderItem *requests.OrderItemRequest, orderID string) error
}
