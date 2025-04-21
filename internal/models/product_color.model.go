package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductColor struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Color     string             `bson:"color,omitempty" json:"color"`
	ImgColor  string             `bson:"img_color,omitempty" json:"img_color"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
