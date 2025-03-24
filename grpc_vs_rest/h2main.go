package main

import (
	"encoding/json"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	_struct "grpc_vs_rest/struct"
	"log"
	"net/http"
)

func h2RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user _struct.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("REST - Received user: %+v", user)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	h2cHandler := h2c.NewHandler(http.HandlerFunc(h2RegisterHandler), &http2.Server{})

	server := &http.Server{
		Addr:    ":8080",
		Handler: h2cHandler,
	}

	log.Println("HTTP/2 (h2c) server is running on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
