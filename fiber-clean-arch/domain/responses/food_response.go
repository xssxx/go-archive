package responses

import "github.com/google/uuid"

type FoodDetail struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Amount   int       `json:"amount"`
	ImageUrl *string   `json:"imageUrl"`	
}
