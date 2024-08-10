package main

import (
	"encoding/json" // JSON ma'lumotlarini kodlash va decoding qilish uchun kutubxona
	"fmt"           // Konsolga xabarlarni yozish uchun kutubxona
	"net/http"      // HTTP server yaratish uchun kutubxona
)

// Message struktura turini aniqlash, JSON formatida "text" maydonini oladi
type Message struct {
	Text string `json:"text"`
}

// jsonHandler funksiyasi HTTP POST so'rovlarini qayta ishlash uchun
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// Agar kiruvchi so'rov POST bo'lmasa, "Method not allowed" xatosini qaytarish
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Kiruvchi JSON so'rovini decoding qilish uchun Message turidagi o'zgaruvchi yaratish
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	// Agar decoding jarayonida xato yuz bersa, "Bad request" xatosini qaytarish
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Kiruvchi so'rovni qayta ishlash va javobni tayyorlash
	response := Message{Text: "Received: " + msg.Text}
	// Javobning Content-Type headerini application/json ga o'rnatish
	w.Header().Set("Content-Type", "application/json")
	// Javobni JSON formatida encoding qilish va mijozga qaytarish
	json.NewEncoder(w).Encode(response)
}

// Main funksiyasi HTTP serverini yaratish va ishga tushirish uchun
func main() {
	// "/json" endpointiga kelgan so'rovlarni jsonHandler funksiyasiga yo'naltirish
	http.HandleFunc("/json", jsonHandler)

	// Serverni ishga tushirish va konsolga xabar yozish
	fmt.Println("server is running http://127.0.0.1:8080/")
	// Serverni port 8080 da ishga tushirish
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// Agar serverni ishga tushirishda xato yuz bersa, xatoni konsolga yozish
		fmt.Println("Failed to start server:", err)
	}
}
