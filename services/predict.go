package services

import (
	"encoding/json"
	"net/http"

	"github.com/heronhoga/auto-biodata-api/models"

)

func Predict(w http.ResponseWriter, r *http.Request) {
	//check if it's application/json
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	prediction := models.Prediction{
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