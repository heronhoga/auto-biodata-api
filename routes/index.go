package routes

import (
	"net/http"

	"github.com/heronhoga/auto-biodata-api/services"
)

func RouteIndex() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /predict", services.Predict)

	return mux
}