package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductSize struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Size      string             `bson:"size,omitempty" json:"size"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
