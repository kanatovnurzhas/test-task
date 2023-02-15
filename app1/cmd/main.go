package main

import (
	"app1/internal/delivery"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/generate-salt", delivery.GenerateSalt)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
