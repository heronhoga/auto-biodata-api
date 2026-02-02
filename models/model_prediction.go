package models

type Prediction struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Nationality string `json:"nationality"`
}

type PredictionResponse struct {
	Status int
	Data Prediction
}