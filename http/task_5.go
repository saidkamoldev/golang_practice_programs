package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Updated: %s", string(body))
}

func main() {
	http.HandleFunc("/update", updateHandler)
	fmt.Println("server is running http://127.0.0.1:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server not running", err)
	}
}
