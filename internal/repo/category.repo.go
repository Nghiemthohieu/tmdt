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

type CategoryRepo struct {
	Collection *mongo.Collection
}

func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{
		Collection: global.Mongodb.Collection("category"),
	}
}

func (r *CategoryRepo) CreateCategory(category *models.Category) (int, error) {
	if category == nil {
		return 20011, errors.New("category cannot be nil")
	}

	_, err := r.Collection.InsertOne(context.TODO(), category)
	if err != nil {
		return 20011, err
	}

	return 20001, nil
}

func (r *CategoryRepo) GetCategoryByID(id string) (*models.Category, int, error) {
	if id == "" {
		return nil, 20014, errors.New("id cannot be empty")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, 20014, errors.New("invalid object id format")
	}

	var category models.Category
	err = r.Collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, 20014, errors.New("no document found with given id")
		}
		return nil, 20014, err
	}

	return &category, 20001, nil
}

func (r *CategoryRepo) GetAllCategories() ([]*models.Category, int, error) {
	var categories []*models.Category
	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, 20015, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var category models.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, 20015, err
		}
		categories = append(categories, &category)
	}

	if err := cursor.Err(); err != nil {
		return nil, 20015, err
	}

	return categories, 20001, nil
}

func (r *CategoryRepo) UpdateCategory(category *models.Category) (int, error) {
	if category == nil {
		return 20012, errors.New("category cannot be nil")
	}

	_, err := r.Collection.UpdateOne(context.TODO(),
		bson.M{"_id": category.ID},
		bson.M{"$set": category},
	)
	if err != nil {
		return 20012, err
	}

	return 20001, nil
}

func (r *CategoryRepo) DeleteCategory(id string) (int, error) {
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
