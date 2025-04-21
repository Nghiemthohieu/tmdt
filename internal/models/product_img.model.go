package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductImage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID primitive.ObjectID `bson:"product_id,omitempty" json:"product_id"` // Tham chiếu đến Product
	ImageURL  string             `bson:"image_url,omitempty" json:"image_url"`
	IsMain    bool               `bson:"is_main,omitempty" json:"is_main"`
	ColorID   primitive.ObjectID `bson:"color_id,omitempty" json:"color_id"` // Tham chiếu đến ProductColor
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
