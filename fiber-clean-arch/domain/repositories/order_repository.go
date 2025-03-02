package repositories

import (
	"context"
	"fiber-clean-arch/domain/models"
	"fiber-clean-arch/domain/requests"
)

type OrderRepository interface {
	Create(ctx context.Context, req *requests.CreateOrderRequest) error
	FindAll(ctx context.Context) ([]models.Order, error)
	FindByFoodID(ctx context.Context) ([]models.Order, error)
	FindByStatus(ctx context.Context) ([]models.Order, error)
}
