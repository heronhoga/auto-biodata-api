package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/heronhoga/auto-biodata-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading env:", err)
	}

	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		log.Fatal("APPLICATION_PORT is not set")
	}

	mux := routes.RouteIndex()

	appPort := fmt.Sprintf(":%s", port)
	fmt.Println("Server is running on port:", port)

	server := &http.Server{
		Addr:         appPort,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
