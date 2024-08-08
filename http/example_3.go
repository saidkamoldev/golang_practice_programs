package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Received:%s", string(body))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func sRunning(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running on 8080")
}

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/status", sRunning)

	fmt.Println("server is running http://127.0.0.1:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
