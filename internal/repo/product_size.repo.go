package repo

import (
	"context"
	"errors"
	"mtb_web/global"
	"mtb_web/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductSizeRepo struct {
	Collection *mongo.Collection
}

func NewProductSizeRepo() *ProductSizeRepo {
	return &ProductSizeRepo{
		Collection: global.Mongodb.Collection("product_size"),
	}
}

func (r *ProductSizeRepo) CreateProductSize(productSize *models.ProductSize) (int, error) {
	if productSize == nil {
		return 20011, errors.New("productSize cannot be nil")
	}

	_, err := r.Collection.InsertOne(context.TODO(), productSize)
	if err != nil {
		return 20011, err
	}

	return 20001, nil
}

func (r *ProductSizeRepo) GetProductSizeByID(id string) (*models.ProductSize, int, error) {
	if id == "" {
		return nil, 20014, errors.New("id cannot be empty")
	}

	var productSize models.ProductSize
	err := r.Collection.FindOne(context.TODO(), map[string]interface{}{"_id": id}).Decode(&productSize)
	if err != nil {
		return nil, 20014, err
	}

	return &productSize, 20001, nil
}

func (r *ProductSizeRepo) GetAllProductSizes() ([]*models.ProductSize, int, error) {
	var productSizes []*models.ProductSize
	cursor, err := r.Collection.Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		return nil, 20015, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var productSize models.ProductSize
		if err := cursor.Decode(&productSize); err != nil {
			return nil, 20015, err
		}
		productSizes = append(productSizes, &productSize)
	}

	if err := cursor.Err(); err != nil {
		return nil, 20015, err
	}

	return productSizes, 20001, nil
}

func (r *ProductSizeRepo) UpdateProductSize(productSize *models.ProductSize) (int, error) {
	if productSize == nil {
		return 20012, errors.New("productSize cannot be nil")
	}

	_, err := r.Collection.UpdateOne(context.TODO(), map[string]interface{}{"_id": productSize.ID}, map[string]interface{}{"$set": productSize})
	if err != nil {
		return 20012, err
	}

	return 20001, nil
}
func (r *ProductSizeRepo) DeleteProductSize(id string) (int, error) {
	if id == "" {
		return 20013, errors.New("id cannot be empty")
	}

	_, err := r.Collection.DeleteOne(context.TODO(), map[string]interface{}{"_id": id})
	if err != nil {
		return 20013, err
	}

	return 20001, nil
}
