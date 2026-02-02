package models

//prediction - general
type PredictionData struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Nationality []Nationality `json:"nationality"`
}

type PredictionRequest struct {
	Name string `json:"name"`
}

type PredictionResponse struct {
	Status int `json:"status"`
	Data PredictionData `json:"data"`
}

//each prediction responses
//agify
type AgifyResponse struct {
	Age int `json:"age"`
}

//genderize
type GenderizeResponse struct {
	Gender string `json:"gender"`
}

//nationalize
type Nationality struct {
	CountryID string `json:"country_id"`
	Probability float64 `json:"probability"`
}
type NationalizeResponse struct {
	Country []Nationality `json:"country"`
}