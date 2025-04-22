package services

import (
	"errors"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductColorService struct {
	ProductColorRepo *repo.ProductColorRepo
}

func NewProductColorService() *ProductColorService {
	return &ProductColorService{
		ProductColorRepo: repo.NewProductColorRepo(),
	}
}

func (s *ProductColorService) CreateProductColor(productColor *models.ProductColor) (int, error) {
	if productColor == nil {
		return 20021, errors.New("productColor cannot be nil")
	}
	newProductColor := &models.ProductColor{
		ID:        primitive.NewObjectID(),
		Color:     productColor.Color,
		ImgColor:  productColor.ImgColor,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	return s.ProductColorRepo.CreateProductColor(newProductColor)
}

func (s *ProductColorService) GetProductColorByID(id string) (*models.ProductColor, int, error) {
	if id == "" {
		return nil, 20024, errors.New("id cannot be empty")
	}
	return s.ProductColorRepo.GetProductColorByID(id)
}

func (s *ProductColorService) GetAllProductColors() ([]*models.ProductColor, int, error) {
	return s.ProductColorRepo.GetAllProductColors()
}

func (s *ProductColorService) UpdateProductColor(productColor *models.ProductColor) (int, error) {
	if productColor == nil {
		return 20022, errors.New("productColor cannot be nil")
	}
	productColor.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return s.ProductColorRepo.UpdateProductColor(productColor)
}

func (s *ProductColorService) DeleteProductColor(id string) (int, error) {
	if id == "" {
		return 20023, errors.New("id cannot be empty")
	}
	return s.ProductColorRepo.DeleteProductColor(id)
}
