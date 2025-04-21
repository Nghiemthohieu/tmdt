package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Code        string             `bson:"code,omitempty" json:"code"`
	Description string             `bson:"description,omitempty" json:"description"`
	Enabled     bool               `bson:"enabled,omitempty" json:"enabled"`
	CreatedAt   primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
