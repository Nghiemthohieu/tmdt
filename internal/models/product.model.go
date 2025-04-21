package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product struct dùng cho MongoDB
type Product struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name          string               `bson:"name,omitempty" json:"name"`
	Description   string               `bson:"description,omitempty" json:"description"`
	Slug          string               `bson:"slug,omitempty" json:"slug"`
	Sizes         []primitive.ObjectID `bson:"sizes,omitempty" json:"sizes"`           // Mối quan hệ many2many
	Colors        []primitive.ObjectID `bson:"colors,omitempty" json:"colors"`         // Mối quan hệ many2many
	Categories    []primitive.ObjectID `bson:"categories,omitempty" json:"categories"` // many2many
	OriginalPrice int                  `bson:"original_price,omitempty" json:"original_price"`
	DiscountPrice int                  `bson:"discount_price,omitempty" json:"discount_price"`
	Price         int                  `bson:"price,omitempty" json:"price"`
	AvgRating     float32              `bson:"avg_rating,omitempty" json:"avg_rating"`
	RatingCount   int                  `bson:"rating_count,omitempty" json:"rating_count"`
	CreatedAt     primitive.DateTime   `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt     primitive.DateTime   `bson:"updated_at,omitempty" json:"updated_at"`
}
