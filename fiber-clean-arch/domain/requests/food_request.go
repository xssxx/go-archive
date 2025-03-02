package requests

type AddMenuRequest struct {
	Name   string  `json:"name" form:"name" validate:"required"`
	Price  float64 `json:"price" form:"price" validate:"required"`
	Amount int     `json:"amount" form:"amount" validate:"required"`
}

type EditMenuRequest struct {
	ID     string  `json:"id" form:"id" validate:"required"`
	Name   string  `json:"name" form:"name" validate:"omitempty"`
	Price  float64 `json:"price" form:"price" validate:"omitempty,gte=0"`
	Amount int     `json:"amount" form:"amount" validate:"omitempty,gte=0"`
}
