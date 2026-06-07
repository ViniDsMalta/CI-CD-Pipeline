package main

import (
	"log"
	"net/http"

	"github.com/ViniDsMalta/CI-CD-Pipeline/internal/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.HealthHandler)

	log.Println("running in 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}