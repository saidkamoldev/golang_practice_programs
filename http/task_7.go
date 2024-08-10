package main

import (
	"fmt"
	"net/http"
)

func delateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Methon not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Resourse deleted")

}

func main() {
	http.HandleFunc("/delete", delateHandler)
	fmt.Println("server is running http://127.0.0.1:8080/")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server is not running", err)
	}
}
