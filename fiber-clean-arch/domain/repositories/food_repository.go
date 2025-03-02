package repositories

import (
	"context"
	"fiber-clean-arch/domain/models"
	"fiber-clean-arch/domain/requests"
)

type FoodRepository interface {
	Create(ctx context.Context, req *requests.AddMenuRequest, imageUrl string) error
	FindAll(ctx context.Context) ([]models.Food, error)
	FindByID(ctx context.Context, id string) (*models.Food, error)
	FindByName(ctx context.Context, name string) (*models.Food, error)
	Edit(ctx context.Context, req *requests.EditMenuRequest, imageUrl string) error
	Delete(ctx context.Context, id string) error
}
