package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name        string               `bson:"name,omitempty" json:"name"`
	Permissions []primitive.ObjectID `bson:"permissions,omitempty" json:"permissions"` // Many-to-many relation, lưu ObjectID của Permissions
	CreatedAt   primitive.DateTime   `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   primitive.DateTime   `bson:"updated_at,omitempty" json:"updated_at"`
}
