package gorm

import (
	"context"
	"fiber-clean-arch/domain/models"
	"fiber-clean-arch/domain/requests"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderGormRepository struct {
	DB *gorm.DB
}

func NewOrderGormRepository(db *gorm.DB) *OrderGormRepository {
	return &OrderGormRepository{
		DB: db,
	}
}

// Create a new order with a timeout context of 5 seconds
func (o *OrderGormRepository) Create(ctx context.Context, req *requests.CreateOrderRequest) error {
	// Set a 5-second timeout for the context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	var orderItems []models.OrderItem
	for _, itemReq := range req.OrderItems {
		foodID, err := uuid.Parse(itemReq.FoodID)
		if err != nil {
			return err
		}
		orderItem := models.OrderItem{
			FoodID:   foodID,
			Quantity: itemReq.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}

	newOrder := models.Order{
		ID:         id,
		Status:     models.PENDING,
		OrderItems: orderItems,
	}

	if err := o.DB.WithContext(ctx).Create(&newOrder).Error; err != nil {
		return err
	}

	return nil
}

// FindAll retrieves all orders with a timeout context of 5 seconds
func (o *OrderGormRepository) FindAll(ctx context.Context) ([]models.Order, error) {
	// Set a 5-second timeout for the context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var orders []models.Order
	if err := o.DB.WithContext(ctx).Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// FindByFoodID retrieves orders that contain a specific food with a timeout context of 5 seconds
func (o *OrderGormRepository) FindByFoodID(ctx context.Context, foodID uuid.UUID) ([]models.Order, error) {
	// Set a 5-second timeout for the context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var orders []models.Order
	if err := o.DB.WithContext(ctx).Preload("OrderItems").Where("id IN (?)", o.DB.Table("order_items").Select("order_id").Where("food_id = ?", foodID)).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// FindByStatus retrieves orders by their status with a timeout context of 5 seconds
func (o *OrderGormRepository) FindByStatus(ctx context.Context, status string) ([]models.Order, error) {
	// Set a 5-second timeout for the context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var orders []models.Order
	if err := o.DB.WithContext(ctx).Preload("OrderItems").Where("status = ?", status).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
