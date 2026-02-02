package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/heronhoga/auto-biodata-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading env: ", err.Error())
	}
	mux := routes.RouteIndex()

	fmt.Println("Server is running on port:", os.Getenv("APPLICATION_PORT"))

	appPort := fmt.Sprintf(":%s", os.Getenv("APPLICATION_PORT"))
	err = http.ListenAndServe(appPort, mux)
	if err != nil {
		log.Println("Error running server: ", err.Error())
	}
}