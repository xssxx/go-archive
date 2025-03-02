package models

import "github.com/google/uuid"

type OrderItem struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrderID  uuid.UUID `json:"orderID" gorm:"type:uuid;not null;index"`
	FoodID   uuid.UUID `json:"foodID" gorm:"type:uuid;not null;index"`
	Quantity int       `json:"quantity" gorm:"type:int;not null;default:1"`

	Order Order `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Food  Food  `gorm:"foreignKey:FoodID;constraint:OnDelete:CASCADE"`
}
