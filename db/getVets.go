package db

import (
	"log"
	"project/go-vets-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetVets() ([]models.Vet, error) {
	col := DB.Collection("vets")

	filter := bson.D{{}}

	var vets []models.Vet
	cur, err := col.Find(Ctx, filter)

	if err != nil {
		log.Println("Error: " + err.Error())
		return vets, err
	}

	err = cur.All(Ctx, &vets)

	if err != nil {
		log.Println("Error: " + err.Error())
		return vets, err
	}

	return vets, nil
}
