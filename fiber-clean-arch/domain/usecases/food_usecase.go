package usecases

import (
	"context"
	"fiber-clean-arch/domain/exceptions"
	"fiber-clean-arch/domain/repositories"
	"fiber-clean-arch/domain/requests"
	"fiber-clean-arch/domain/responses"
	"fiber-clean-arch/internal/adapters/storage"
	"mime/multipart"
)

type FoodUseCase interface {
	Create(ctx context.Context, req *requests.AddMenuRequest, file multipart.File) error
	FindAll(ctx context.Context) ([]responses.FoodDetail, error)
	FindByID(ctx context.Context, id string) (*responses.FoodDetail, error)
	FindByName(ctx context.Context, name string) (*responses.FoodDetail, error)
	Edit(ctx context.Context, req *requests.EditMenuRequest, file multipart.File) error
	Delete(ctx context.Context, id string) error
}

type FoodService struct {
	foodRepo       repositories.FoodRepository
	storageService storage.StorageService
}

func NewFoodService(foodRepo repositories.FoodRepository, storage storage.StorageService) FoodUseCase {
	return &FoodService{
		foodRepo:       foodRepo,
		storageService: storage,
	}
}

func (f *FoodService) Create(ctx context.Context, req *requests.AddMenuRequest, file multipart.File) error {
	food, err := f.foodRepo.FindByName(ctx, req.Name)
	if err != nil {
		return err
	}
	if food != nil {
		return exceptions.ErrDuplicatedFoodName
	}

	imageUrl, err := f.storageService.UploadFile(ctx, file)
	if err != nil {
		return err
	}

	return f.foodRepo.Create(ctx, req, imageUrl)
}

func (f *FoodService) FindAll(ctx context.Context) ([]responses.FoodDetail, error) {
	foods, err := f.foodRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	foodDetails := make([]responses.FoodDetail, 0)

	for _, food := range foods {
		foodDetails = append(foodDetails, responses.FoodDetail{
			ID:       food.ID,
			Name:     food.Name,
			Price:    food.Price,
			Amount:   food.Amount,
			ImageUrl: food.ImageUrl,
		})
	}

	return foodDetails, nil
}

func (f *FoodService) FindByID(ctx context.Context, id string) (*responses.FoodDetail, error) {
	food, err := f.foodRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	foodDetail := &responses.FoodDetail{
		ID:       food.ID,
		Name:     food.Name,
		Price:    food.Price,
		Amount:   food.Amount,
		ImageUrl: food.ImageUrl,
	}
	return foodDetail, nil
}

func (f *FoodService) FindByName(ctx context.Context, name string) (*responses.FoodDetail, error) {
	food, err := f.foodRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	foodDetail := &responses.FoodDetail{
		ID:       food.ID,
		Name:     food.Name,
		Price:    food.Price,
		Amount:   food.Amount,
		ImageUrl: food.ImageUrl,
	}
	return foodDetail, nil
}

func (f *FoodService) Edit(ctx context.Context, req *requests.EditMenuRequest, file multipart.File) error {
	food, err := f.foodRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if food == nil {
		return exceptions.ErrFoodNotFound
	}
	if file == nil {
		return f.foodRepo.Edit(ctx, req, *food.ImageUrl)
	}

	imageUrl, err := f.storageService.UploadFile(ctx, file)
	if err != nil {
		return err
	}

	return f.foodRepo.Edit(ctx, req, imageUrl)
}

func (f *FoodService) Delete(ctx context.Context, id string) error {
	food, err := f.foodRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if food == nil {
		return exceptions.ErrFoodNotFound
	}

	return f.foodRepo.Delete(ctx, id)
}
