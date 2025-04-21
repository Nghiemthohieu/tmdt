package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductReview struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID primitive.ObjectID `bson:"product_id,omitempty" json:"product_id"` // Tham chiếu tới Product
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`       // Tham chiếu tới User
	Rating    int                `bson:"rating,omitempty" json:"rating"`
	Review    string             `bson:"review,omitempty" json:"review"`
	Images    []string           `bson:"images,omitempty" json:"images"` // Sử dụng array string thay vì table riêng
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
