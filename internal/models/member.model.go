package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	Phone     string             `bson:"phone,omitempty" json:"phone"`
	Birthdate time.Time          `bson:"birthdate,omitempty" json:"birthdate"` // Sử dụng time.Time cho date
	Address   string             `bson:"address,omitempty" json:"address"`
	Avatar    string             `bson:"avatar,omitempty" json:"avatar"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
