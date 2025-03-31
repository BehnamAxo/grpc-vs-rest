package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type HistoryEntry struct {
	Year  int    `json:"year"`
	Event string `json:"event"`
}

type User struct {
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	Bio       string         `json:"bio"`
	Interests []string       `json:"interests"`
	History   []HistoryEntry `json:"history"`
}

var requestCounter uint64

func processUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	count := atomic.AddUint64(&requestCounter, 1)

	var user User

	// Limiting request body size to 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("❌ [%d] JSON decode failed: %v\n", count, err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("[%d] Received: %s | Interests: %d | History: %d | Bio Size: %d bytes\n",
		count, user.Name, len(user.Interests), len(user.History), len(user.Bio))

	resp := map[string]string{
		"message": "Payment processed for " + user.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/process", processUserHandler)
	fmt.Println("REST payment service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
