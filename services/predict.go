package services

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/heronhoga/auto-biodata-api/models"
)

func Predict(w http.ResponseWriter, r *http.Request) {
	//headers check - content type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	//headers check - application key
	appKey := r.Header.Get("App-Key")
	if appKey != os.Getenv("APPLICATION_KEY") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//check request body
	if r.ContentLength == 0 {
		http.Error(w, "payload is required", http.StatusBadRequest)
		return
	}

	var request models.PredictionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if request.Name == "" {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}
	
	//service
	prediction := models.PredictionData{
		Name: "Hoga",
		Age: 20,
		Gender: "Male",
		Nationality: "ID",
	}

	response := models.PredictionResponse{
		Status: 200,
		Data: prediction,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}