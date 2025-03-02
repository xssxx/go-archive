package models

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	PENDING    Status = "peding"
	PAID       Status = "paid"
	COOKING    Status = "cooking"
	DELIVERING Status = "delivering"
	SUCCESS    Status = "success"
	CANCELLED  Status = "cancelled"
)

type Order struct {
	ID         uuid.UUID   `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Status     Status      `json:"status" gorm:"type:varchar(20);not null;default:'pending'"`
	OrderItems []OrderItem `json:"orderItems" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
}
