package services

import (
	"errors"
	"fmt"
	"mtb_web/internal/models"
	"mtb_web/internal/repo"
	util "mtb_web/pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService struct {
	CategoryRepo *repo.CategoryRepo
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		CategoryRepo: repo.NewCategoryRepo(),
	}
}

func (s *CategoryService) CreateCategory(category *models.Category) (int, error) {
	if category == nil {
		return 20031, errors.New("category cannot be nil")
	}
	imgData, err := util.DecodeBase64Image(category.Img)
	if err != nil {
		return 20031, err
	}

	imgURL, err := util.UpLoadFile(imgData)
	if err != nil {
		return 20031, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	category.Img = imgURL
	newCategory := &models.Category{
		ID:         primitive.NewObjectID(),
		Name:       category.Name,
		Slug:       category.Slug,
		Img:        category.Img,
		ParentID:   category.ParentID,
		ProductIDs: category.ProductIDs,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}
	return s.CategoryRepo.CreateCategory(newCategory)
}

func (s *CategoryService) GetCategoryByID(id string) (*models.Category, int, error) {
	if id == "" {
		return nil, 20034, errors.New("id cannot be empty")
	}
	return s.CategoryRepo.GetCategoryByID(id)
}

func (s *CategoryService) GetAllCategories() ([]*models.Category, int, error) {
	return s.CategoryRepo.GetAllCategories()
}

func (s *CategoryService) UpdateCategory(category *models.Category) (int, error) {
	if category == nil {
		return 20032, errors.New("category cannot be nil")
	}
	imgData, err := util.DecodeBase64Image(category.Img)
	if err != nil {
		return 20032, err
	}

	imgURL, err := util.UpLoadFile(imgData)
	if err != nil {
		return 20032, fmt.Errorf("failed to upload image to S3: %v", err)
	}
	category.Img = imgURL
	category.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return s.CategoryRepo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(id string) (int, error) {
	if id == "" {
		return 20033, errors.New("id cannot be empty")
	}
	return s.CategoryRepo.DeleteCategory(id)
}
