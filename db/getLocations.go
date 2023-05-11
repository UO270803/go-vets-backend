package db

import (
	"log"
	"project/go-vets-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetLocations() ([]models.Location, error) {
	col := DB.Collection("locations")

	filter := bson.D{{}}

	var locations []models.Location
	cur, err := col.Find(Ctx, filter)

	if err != nil {
		log.Println("Error: " + err.Error())
		return locations, err
	}

	err = cur.All(Ctx, &locations)

	if err != nil {
		log.Println("Error: " + err.Error())
		return locations, err
	}

	return locations, nil
}
