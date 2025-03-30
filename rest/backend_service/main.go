package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	user := User{
		Name:  "Sir Laughsalot McGiggles",
		Age:   420,
		Email: "funny.bone@laughterverse.io",
		Phone: "+1-800-GIGGLEZ",
	}

	body, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8080/process", "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("🔥Error sending request: %v", err)
	}

	defer resp.Body.Close()

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)

	fmt.Println("Backend received response:", response["message"])
}
