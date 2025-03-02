package exceptions

import "errors"

var (
	ErrDuplicatedFoodName = errors.New("food name already exist")
	ErrFoodNotFound       = errors.New("food not found")
)
