package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Email     string               `bson:"email,omitempty" json:"email"`
	Password  string               `bson:"password,omitempty" json:"password"`
	Role      []primitive.ObjectID `bson:"role,omitempty" json:"role"` // Many-to-many relation, lưu ObjectID
	CreatedAt primitive.DateTime   `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime   `bson:"updated_at,omitempty" json:"updated_at"`
}
