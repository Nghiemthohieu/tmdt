package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name,omitempty" json:"name"`
	UserID          primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"` // Tham chiếu tới User
	Email           string             `bson:"email,omitempty" json:"email"`
	Phone           string             `bson:"phone,omitempty" json:"phone"`
	Address         string             `bson:"address,omitempty" json:"address"`
	TotalPrice      float64            `bson:"total_price,omitempty" json:"total_price"`
	PaymentMethodID primitive.ObjectID `bson:"payment_method_id,omitempty" json:"payment_method_id"` // Tham chiếu tới PaymentMethod
	Status          string             `bson:"status,omitempty" json:"status"`
	CreatedAt       primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt       primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
