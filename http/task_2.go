package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Methot not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Qabul qilindi: %s", string(body))

}

func main() {
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("server is running http://127.0.0.1:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Fail tpo start", err)
	}
}
