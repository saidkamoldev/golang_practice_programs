package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func sRunning(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running on 8080")
}

func main() {

	http.HandleFunc("/status", sRunning)
	http.HandleFunc("/hello", helloHandler) // URL yo'lini va uni qayta ishlovchi funcksiyani belgilash

	fmt.Println("server is running http://127.0.0.1:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
