package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permissions struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Method      string             `bson:"method,omitempty" json:"method"`
	Path        string             `bson:"path,omitempty" json:"path"`
	Description string             `bson:"description,omitempty" json:"description"`
	Module      string             `bson:"module,omitempty" json:"module"`
	CreatedAt   primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
}
