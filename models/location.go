package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	City bson.M             `bson:"city" json:"city"`
}
