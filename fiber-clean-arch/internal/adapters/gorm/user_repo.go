package gorm

import (
	"context"
	"errors"
	"fiber-clean-arch/domain/models"
	"fiber-clean-arch/domain/requests"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	DB *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{
		DB: db,
	}
}

func (u *UserGormRepository) Create(ctx context.Context, req *requests.UserRegisterRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	newUser := models.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     models.CUSTOMER,
	}

	if err := u.DB.WithContext(ctx).Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserGormRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user models.User
	if err := u.DB.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}

	return &user, nil
}
