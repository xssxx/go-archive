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

type OrderItemGormRepository struct {
	DB *gorm.DB
}

func NewOrderItemGormRepository(db *gorm.DB) *OrderItemGormRepository {
	return &OrderItemGormRepository{
		DB: db,
	}
}

// GetOrderItemsByOrderID retrieves order items by the associated order ID
func (o *OrderItemGormRepository) GetOrderItemsByOrderID(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	// Set a 5-second timeout for the context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var orderItems []models.OrderItem
	if err := o.DB.WithContext(ctx).Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if no order items are found
		}
		return nil, err
	}

	return orderItems, nil
}

// CreateOrderItems creates a new order item for a given order ID
func (o *OrderItemGormRepository) CreateOrderItems(ctx context.Context, req *requests.OrderItemRequest, orderID string) error {
	// Set a 5-second timeout for the context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Parse the FoodID into a uuid
	foodID, err := uuid.Parse(req.FoodID)
	if err != nil {
		return err
	}

	// Create the order item
	orderItem := models.OrderItem{
		ID:       uuid.New(),
		OrderID:  uuid.MustParse(orderID), // Assuming orderID is a valid uuid string
		FoodID:   foodID,
		Quantity: req.Quantity,
	}

	// Insert the order item into the database
	if err := o.DB.WithContext(ctx).Create(&orderItem).Error; err != nil {
		return err
	}

	return nil
}
