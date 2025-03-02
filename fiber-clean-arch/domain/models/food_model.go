package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Food struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Price     float64        `json:"price" gorm:"type:decimal(10,2);not null"`
	Amount    int            `json:"amount" gorm:"not null;default:0"`
	ImageUrl  *string        `json:"imageUrl" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
