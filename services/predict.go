package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/heronhoga/auto-biodata-api/models"
	"github.com/heronhoga/auto-biodata-api/utils"
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

	var request models.PredictionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		if err == io.EOF {
			http.Error(w, "payload cannot be empty", http.StatusBadRequest)
			return
		}
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if request.Name == "" {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}
	
	//service call
	agifyUrl := os.Getenv("AGIFY_URL")
	genderizeUrl := os.Getenv("GENDERIZE_URL")
	nationalizeUrl := os.Getenv("NATIONALIZE_URL")

	ch := make(chan models.ApiResult)
	var wg sync.WaitGroup
	wg.Add(3)

	go utils.Fetch("agify", agifyUrl+request.Name, ch, &wg)
	go utils.Fetch("genderize", genderizeUrl+request.Name, ch, &wg)
	go utils.Fetch("nationalize", nationalizeUrl+request.Name, ch, &wg)
	
	go func ()  {
		wg.Wait()
		close(ch)
	}()

	var agifyResponse models.AgifyResponse
	var genderizeResponse models.GenderizeResponse
	var nationalizeResponse models.NationalizeResponse

	for result := range ch {
		if result.Err != nil {
			log.Println(result.Service, "error:", result.Err)
			continue
		}

		switch result.Service {
		case "agify":
			if err := json.Unmarshal(result.Body, &agifyResponse); err != nil {
				log.Println("agify unmarshal error:", err)
			}

		case "genderize":
			if err := json.Unmarshal(result.Body, &genderizeResponse); err != nil {
				log.Println("genderize unmarshal error:", err)
			}

		case "nationalize":
			if err := json.Unmarshal(result.Body, &nationalizeResponse); err != nil {
				log.Println("nationalize unmarshal error:", err)
			}
		}
	}

	//check data
	fmt.Println(agifyResponse)
	fmt.Println(genderizeResponse)
	fmt.Println(nationalizeResponse)
	
	//end service call
	response := models.PredictionResponse{
		Status: 200,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}