package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductVariant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID primitive.ObjectID `bson:"product_id,omitempty" json:"product_id"` // Tham chiếu đến Product
	SizeID    primitive.ObjectID `bson:"size_id,omitempty" json:"size_id"`       // Tham chiếu đến ProductSize
	ColorID   primitive.ObjectID `bson:"color_id,omitempty" json:"color_id"`     // Tham chiếu đến ProductColor
	Stock     int                `bson:"stock,omitempty" json:"stock"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
