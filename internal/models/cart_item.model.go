package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItems struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CartID    primitive.ObjectID `bson:"cart_id,omitempty" json:"cart_id"`
	ProductID primitive.ObjectID `bson:"product_id,omitempty" json:"product_id"`
	SizeID    primitive.ObjectID `bson:"size_id,omitempty" json:"size_id"`
	ColorID   primitive.ObjectID `bson:"color_id,omitempty" json:"color_id"`
	Quantity  int                `bson:"quantity,omitempty" json:"quantity"`
}
