package main

import (
	"fmt"
	"net/http"
)

func oldHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/new", http.StatusMovedPermanently)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the new location!"))
}

func main() {
	http.HandleFunc("/old", oldHandler)
	http.HandleFunc("/new", newHandler)
	fmt.Println("server is running http://127.0.0.1:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server not running on :8080", err)
	}
}
