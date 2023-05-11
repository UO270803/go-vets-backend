package routers

import (
	"encoding/json"
	"net/http"
	"project/go-vets-backend/db"
)

func GetVets(response http.ResponseWriter, request *http.Request) {

	vets, err := db.GetVets()
	if err != nil {
		http.Error(response, "Error trying to find the record: "+err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("context-type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(vets)
}
