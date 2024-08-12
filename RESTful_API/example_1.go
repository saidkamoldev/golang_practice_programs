package main

import (
	"encoding/json" // JSON formatida kodlash va dekodlash uchun ishlatiladi
	"fmt"           // Konsolga yozish uchun ishlatiladi
	"io/ioutil"
	"net/http" // HTTP protokoli bilan ishlash uchun paket
	"strconv"  // Stringni integerga o'girish uchun ishlatiladi

	"github.com/gorilla/mux" // Gorilla/Mux paketini import qiladi, bu routerni boshqarish uchun ishlatiladi
)

// User strukturasini aniqlash
type User struct {
	ID   int    `json:"id"`   // ID maydoni JSON da "id" deb nomlanadi
	Name string `json:"Name"` // Name maydoni JSON da "Name" deb nomlanadi
}

type User2 struct {
	ID   int    `json:"id"`   // ID maydoni JSON da "id" deb nomlanadi
	Name string `json:"Name"` // Name maydoni JSON da "Name" deb nomlanadi
}

// Foydalanuvchilar ro'yxati, dastlab 2 ta foydalanuvchi bilan to'ldirilgan
var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

var users2 = []User2{
	{ID: 1, Name: "TEST_NAME"},
	{ID: 2, Name: "TEST_NAME2"},
}

// getUsers funksiyasi barcha foydalanuvchilarni JSON formatida qaytaradi
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // Javobni JSON formatida qaytarish uchun sarlavha qo'shadi
	json.NewEncoder(w).Encode(users)                   // Foydalanuvchilarni JSON formatida kodlaydi va javob sifatida yuboradi
	json.NewEncoder(w).Encode(users2)
}

// getUser funksiyasi URL dan kelgan ID bo'yicha bitta foydalanuvchini qaytaradi
func getUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var result []User

	for _, user := range users {
		if user.ID == id {
			result = append(result, user)
		}
	}

	// ikkinchi

	for _, user := range users2 {
		if user.ID == id {
			result = append(result, User(user))
		}
	}

	if len(result) == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

	//  1 ta structli funksiya
	// params := mux.Vars(r)               // URL dan parametrlarni olish (shu jumladan "id")
	// id, _ := strconv.Atoi(params["id"]) // ID ni integer formatiga o'zgartirish

	// for _, user := range users { // Barcha foydalanuvchilarni ko'rib chiqish
	// 	if user.ID == id { // Agar foydalanuvchi ID mos kelsa
	// 		w.Header().Set("Content-Type", "application/json") // Javob sarlavhasini JSON formatiga sozlash
	// 		json.NewEncoder(w).Encode(user)                    // Foydalanuvchini JSON formatida kodlaydi va javob sifatida yuboradi
	// 		return                                             // Funksiyani tugatadi (keyingi kod bajarilmaydi)
	// 	}
	// }

	// http.Error(w, "User not found", http.StatusNotFound) // Agar foydalanuvchi topilmasa, 404 xato javobi yuboriladi
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)               // URL dan parametrlarni olish (shu jumladan "id")
	id, _ := strconv.Atoi(params["id"]) // ID ni integer formatiga o'zgartirish

	var delete bool

	for index, user := range users { // Barcha foydalanuvchilarni ko'rib chiqish
		if user.ID == id { // Agar foydalanuvchi ID mos kelsa
			users = append(users[:index], users[index+1:]...) // Ro'yxatdan foydalanuvchini o'chirish (index bo'yicha)
			delete = true
			fmt.Fprintf(w, "User deleted from users \n")
			w.WriteHeader(http.StatusNoContent) // 204 status kodi bilan muvaffaqiyatli javob qaytarish (kontent yo'q)
			// Funksiyani tugatadi
		}
	}

	for index, user := range users2 {
		if user.ID == id {
			users2 = append(users2[:index], users2[index+1:]...)
			fmt.Fprintf(w, "User deleted from users2 ")
			delete = true
		}
	}

	if delete == true {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, "User not found", http.StatusNotFound) // Agar foydalanuvchi topilmasa, 404 xato javobi yuboriladi
}

// createUser funksiyasi yangi foydalanuvchini yaratadi va uni foydalanuvchilar ro'yxatiga qo'shadi
func createUser(w http.ResponseWriter, r *http.Request) {

	var newUser User
	var newUser2 User2

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	err = json.Unmarshal(body, &newUser2)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	newUser2.ID = len(users2) + 1
	users2 = append(users2, newUser2)

	w.Header().Set("Content-Type", "application/json")
	// ikkila yangi foydalanuvchini javob sifatida qaytaradi
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user1": newUser,
		"user2": newUser2,
	})

	// bitta struckli
	// var newUser User // Yangi foydalanuvchi uchun o'zgaruvchi yaratish
	// json.NewDecoder(r.Body).Decode(&newUser) // HTTP so'rovining tanasidan (body) JSON formatidagi ma'lumotni o'qib olish va newUser ga o'zgartirish
	// newUser.ID = len(users) + 1              // Yangi foydalanuvchiga ID tayinlash (foydalanuvchilar soniga +1)
	// users = append(users, newUser)           // Yangi foydalanuvchini ro'yxatga qo'shish

	// w.Header().Set("Content-Type", "application/json") // Javob sarlavhasini JSON formatiga sozlash
	// json.NewEncoder(w).Encode(newUser)                 // Yangi foydalanuvchini JSON formatida kodlaydi va javob sifatida yuboradi
}

// deleteUser funksiyasi URL dan kelgan ID bo'yicha foydalanuvchini o'chiradi

// Asosiy funksiya
func main() {
	r := mux.NewRouter() // Yangi router yaratish (mux paketi yordamida)

	// URL yo'llari va ularni qayta ishlovchi funksiyalarni aniqlash
	r.HandleFunc("/users", getUsers).Methods("GET")           // Barcha foydalanuvchilarni olish
	r.HandleFunc("/users/{id}", getUser).Methods("GET")       // Bitta foydalanuvchini olish
	r.HandleFunc("/users", createUser).Methods("POST")        // Yangi foydalanuvchi yaratish
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE") // Foydalanuvchini o'chirish

	fmt.Print("Server is running on port :8080") // Konsolga server ishga tushganligi haqida xabar berish

	// Serverni 8080 portda ishga tushirish
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Server not running:", err) // Agar serverni ishga tushirishda xato yuzaga kelsa, xabarni chop etish
	}
}
