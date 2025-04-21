package services

import (
	"errors"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductSizeService struct {
	ProductSizeRepo *repo.ProductSizeRepo
}

func NewProductSizeService() *ProductSizeService {
	return &ProductSizeService{
		ProductSizeRepo: repo.NewProductSizeRepo(),
	}
}

func (s *ProductSizeService) CreateProductSize(productSize *models.ProductSize) (int, error) {
	if productSize == nil {
		return 20011, errors.New("productSize cannot be nil")
	}
	NewproductSize := &models.ProductSize{
		ID:        primitive.NewObjectID(),
		Size:      productSize.Size,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	return s.ProductSizeRepo.CreateProductSize(NewproductSize)
}

func (s *ProductSizeService) GetProductSizeByID(id string) (*models.ProductSize, int, error) {
	if id == "" {
		return nil, 20014, errors.New("id cannot be empty")
	}
	return s.ProductSizeRepo.GetProductSizeByID(id)
}

func (s *ProductSizeService) GetAllProductSizes() ([]*models.ProductSize, int, error) {
	return s.ProductSizeRepo.GetAllProductSizes()
}

func (s *ProductSizeService) UpdateProductSize(productSize *models.ProductSize) (int, error) {
	if productSize == nil {
		return 20012, errors.New("productSize cannot be nil")
	}
	productSize.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return s.ProductSizeRepo.UpdateProductSize(productSize)
}

func (s *ProductSizeService) DeleteProductSize(id string) (int, error) {
	if id == "" {
		return 20013, errors.New("id cannot be empty")
	}
	return s.ProductSizeRepo.DeleteProductSize(id)
}
