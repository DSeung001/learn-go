package main

import (
	"encoding/json"
	_struct "grpc_vs_rest/struct"
	"log"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user _struct.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 실제 회원가입 로직 (예: DB 저장)
		log.Printf("REST - Received user: %+v", user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/register", registerHandler)
	log.Println("REST API server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
