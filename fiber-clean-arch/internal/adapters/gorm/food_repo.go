package gorm

import (
	"context"
	"errors"
	"fiber-clean-arch/domain/models"
	"fiber-clean-arch/domain/requests"
	"time"

	"gorm.io/gorm"
)

type FoodGormRepository struct {
	DB *gorm.DB
}

func NewFoodGormRepository(db *gorm.DB) *FoodGormRepository {
	return &FoodGormRepository{
		DB: db,
	}
}

// Create a new food entry in the database
func (f *FoodGormRepository) Create(ctx context.Context, req *requests.AddMenuRequest, imageUrl string) error {
	// Set context timeout to 5 seconds (you can adjust this based on your needs)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	food := models.Food{
		Name:     req.Name,
		Price:    req.Price,
		Amount:   req.Amount,
		ImageUrl: &imageUrl, // Assuming imageUrl is optional and nullable
	}

	if err := f.DB.WithContext(ctx).Create(&food).Error; err != nil {
		return err
	}

	return nil
}

// FindAll retrieves all food entries from the database
func (f *FoodGormRepository) FindAll(ctx context.Context) ([]models.Food, error) {
	// Set context timeout to 5 seconds
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var foods []models.Food
	if err := f.DB.WithContext(ctx).Find(&foods).Error; err != nil {
		return nil, err
	}

	return foods, nil
}

// FindByID retrieves a single food entry by its ID
func (f *FoodGormRepository) FindByID(ctx context.Context, id string) (*models.Food, error) {
	// Set context timeout to 5 seconds
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var food models.Food
	if err := f.DB.WithContext(ctx).First(&food, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}

	return &food, nil
}

// FindByName retrieves a food entry by its name
func (f *FoodGormRepository) FindByName(ctx context.Context, name string) (*models.Food, error) {
	// Set context timeout to 5 seconds
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var food models.Food
	if err := f.DB.WithContext(ctx).First(&food, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}

	return &food, nil
}

// Edit modifies an existing food entry
func (f *FoodGormRepository) Edit(ctx context.Context, req *requests.EditMenuRequest, imageUrl string) error {
	// Set context timeout to 5 seconds
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var food models.Food
	// Find the food entry by ID
	if err := f.DB.WithContext(ctx).First(&food, "id = ?", req.ID).Error; err != nil {
		return err // Return error if the food does not exist
	}

	// Update the fields of the food entry
	food.Name = req.Name
	food.Price = req.Price
	food.Amount = req.Amount
	if imageUrl != "" {
		food.ImageUrl = &imageUrl
	}

	// Save the updated food entry
	if err := f.DB.WithContext(ctx).Save(&food).Error; err != nil {
		return err
	}

	return nil
}

// Delete removes a food entry by its ID
func (f *FoodGormRepository) Delete(ctx context.Context, id string) error {
	// Set context timeout to 5 seconds
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var food models.Food
	// Find the food entry by ID
	if err := f.DB.WithContext(ctx).First(&food, "id = ?", id).Error; err != nil {
		return err // Return error if the food does not exist
	}

	// Delete the food entry
	if err := f.DB.WithContext(ctx).Delete(&food).Error; err != nil {
		return err
	}

	return nil
}
