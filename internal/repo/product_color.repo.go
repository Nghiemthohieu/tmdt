package repo

import (
	"context"
	"errors"
	"mtb_web/global"
	"mtb_web/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductColorRepo struct {
	Collection *mongo.Collection
}

func NewProductColorRepo() *ProductColorRepo {
	return &ProductColorRepo{
		Collection: global.Mongodb.Collection("product_color"),
	}
}

func (r *ProductColorRepo) CreateProductColor(productColor *models.ProductColor) (int, error) {
	if productColor == nil {
		return 20011, errors.New("productColor cannot be nil")
	}

	_, err := r.Collection.InsertOne(context.TODO(), productColor)
	if err != nil {
		return 20011, err
	}

	return 20001, nil
}

func (r *ProductColorRepo) GetProductColorByID(id string) (*models.ProductColor, int, error) {
	if id == "" {
		return nil, 20014, errors.New("id cannot be empty")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, 20014, errors.New("invalid object id format")
	}

	var productColor models.ProductColor
	err = r.Collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&productColor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, 20014, errors.New("no document found with given id")
		}
		return nil, 20014, err
	}

	return &productColor, 20001, nil
}

func (r *ProductColorRepo) GetAllProductColors() ([]*models.ProductColor, int, error) {
	var productColors []*models.ProductColor
	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, 20015, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var productColor models.ProductColor
		if err := cursor.Decode(&productColor); err != nil {
			return nil, 20015, err
		}
		productColors = append(productColors, &productColor)
	}

	if err := cursor.Err(); err != nil {
		return nil, 20015, err
	}

	return productColors, 20001, nil
}

func (r *ProductColorRepo) UpdateProductColor(productColor *models.ProductColor) (int, error) {
	if productColor == nil {
		return 20012, errors.New("productColor cannot be nil")
	}

	_, err := r.Collection.UpdateOne(context.TODO(),
		bson.M{"_id": productColor.ID},
		bson.M{"$set": productColor},
	)
	if err != nil {
		return 20012, err
	}

	return 20001, nil
}

func (r *ProductColorRepo) DeleteProductColor(id string) (int, error) {
	if id == "" {
		return 20013, errors.New("id cannot be empty")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 20013, errors.New("invalid object id format")
	}

	res, err := r.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return 20013, err
	}
	if res.DeletedCount == 0 {
		return 20014, errors.New("no document found to delete")
	}

	return 20001, nil
}
