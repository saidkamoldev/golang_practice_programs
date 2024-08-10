package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("Name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	fmt.Println("server is running http://127.0.0.1:8080/")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server is not running", err)
	}
}
