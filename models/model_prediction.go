package models

type PredictionData struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Nationality string `json:"nationality"`
}

type PredictionRequest struct {
	Name string `json:"name"`
}

type PredictionResponse struct {
	Status int
	Data PredictionData
}