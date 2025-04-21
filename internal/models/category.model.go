package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string               `bson:"name,omitempty" json:"name,omitempty"`
	Slug       string               `bson:"slug,omitempty" json:"slug,omitempty"`
	Img        string               `bson:"img,omitempty" json:"img,omitempty"`
	ParentID   primitive.ObjectID   `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
	ProductIDs []primitive.ObjectID `bson:"product_ids,omitempty" json:"product_ids,omitempty"`
	CreatedAt  primitive.DateTime   `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
