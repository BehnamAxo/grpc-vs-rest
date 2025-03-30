package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var requestCounter uint64

func processUserHandler(w http.ResponseWriter, r *http.Request) {
	count := atomic.AddUint64(&requestCounter, 1)

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("🔥[%d] Failed to decode JSON: %v\n", count, err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("[%d] Payment service received user: %s (%s)\n", count, user.Name, user.Email)

	resp := map[string]string{"message": "Payment processed for " + user.Name}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/process", processUserHandler)
	fmt.Println("REST payment service running on port 8080!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// autocannon -c 100 -d 20 -m POST -H "Content-Type: application/json" -b "{\"name\":\"Sir Laughsalot McGiggles\",\"age\":420,\"email\":\"funny.bone@laughterverse.io\",\"phone\":\"+1-800-GIGGLEZ\"}" http://localhost:8080/process
