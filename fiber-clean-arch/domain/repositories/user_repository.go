package repositories

import (
	"context"
	"fiber-clean-arch/domain/models"
	"fiber-clean-arch/domain/requests"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegisterRequest) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
