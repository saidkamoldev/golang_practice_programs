package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Metod not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello, Golang")

}

func main() {

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("server is running http://127.0.0.1:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Faild to start server:", err)
	}

	// fmt.Println("server is running http://127.0.0.1:8080/")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("Failed to start server:", err)
	// }
}
