package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vet struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Coord string             `bson:"coord" json:"coord"`
	City  bson.M             `bson:"city,omitempty" json:"city,omitempty"`
}
