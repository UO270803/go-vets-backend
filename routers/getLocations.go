package routers

import (
	"encoding/json"
	"net/http"
	"project/go-vets-backend/db"
)

func GetLocations(response http.ResponseWriter, request *http.Request) {

	locations, err := db.GetLocations()
	if err != nil {
		http.Error(response, "Error trying to find the record: "+err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("context-type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(locations)
}
