package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Staffs struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Phone     string             `bson:"phone,omitempty" json:"phone"`
	Position  string             `bson:"position,omitempty" json:"position"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
