package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheckEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is alive.")
	fmt.Println("Status: Healthy")
}

func handleRequests() {
	http.HandleFunc("/", healthCheckEndpoint)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}

