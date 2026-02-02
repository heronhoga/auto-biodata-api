package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/heronhoga/auto-biodata-api/routes"
)

func main() {
	mux := routes.RouteIndex()

	fmt.Println("Server is running")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Error running server: ", err.Error())
	}
}