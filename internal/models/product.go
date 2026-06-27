package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Product struct {
	ID bson.ObjectID	 `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name" binding:"required"`
	Price int `bson:"price" json:"price" binding:"required,gt=0"`
	SellerID int `bson:"seller_id,omitempty" json:"seller_id"`
	Stock int `bson:"stock" json:"stock" binding:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}