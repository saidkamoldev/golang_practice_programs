package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Methot not allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := Message{Text: "Message received:" + msg.Text}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// http.HandlerFunc("/message", messageHandler)
	http.HandleFunc("/message", messageHandler)
	fmt.Println("server is running http://127.0.0.1:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server not running", err)
	}
}
