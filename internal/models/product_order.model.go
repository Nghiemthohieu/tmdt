package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductOrder struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID   primitive.ObjectID `bson:"order_id,omitempty" json:"order_id"`     // Reference đến Order
	ProductID primitive.ObjectID `bson:"product_id,omitempty" json:"product_id"` // Reference đến Product
	SizeID    primitive.ObjectID `bson:"size_id,omitempty" json:"size_id"`       // Reference đến ProductSize
	ColorID   primitive.ObjectID `bson:"color_id,omitempty" json:"color_id"`     // Reference đến ProductColor
	Quantity  int                `bson:"quantity,omitempty" json:"quantity"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
