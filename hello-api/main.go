package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type greeting struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	msg := "Hello, World!"
	if name != "" {
		msg = fmt.Sprintf("Hello, %s!", name)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(greeting{Message: msg}); err != nil {
		log.Printf("error encoding response: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /health", healthHandler)

	log.Println("listening on :9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}
