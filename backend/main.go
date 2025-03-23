package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

// enableCors adds CORS headers to responses
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CORS middleware triggered for:", r.Method, r.URL.Path)
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return // Stop further processing for preflight requests
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request for /api/hello")
		response := Response{Message: "Hello from Go backend!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	fmt.Println("Go server running on http://localhost:4000")
	http.ListenAndServe(":4000", enableCors(r))
}
