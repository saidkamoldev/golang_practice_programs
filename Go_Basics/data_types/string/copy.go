package main

import "fmt"

func main() {
	// string yaratish
	var hello string = "Salom, Dunyo!"

	// Stringni elon qilish va qiymat berish
	greeting := "Qanday ahvolda"

	// Stringni chop etish
	fmt.Println(hello)
	fmt.Println(greeting)

	// Stringni uzunligini topish
	fmt.Println("Hello string uzunligi:", len(hello))

	//stringni belgilarini olish
	// fmt.Println("Birinchi belgini:", string(hello[0]))

	// Birinchi uslub ancha tushunarli ammo UTF-8 bilan muammo yuzaga keladi misol uchun "Привет"da no'to'g'ri ishlaydi;
	// C uslubiga o'xshash bo'lganligi sababli, dasturchilar uchun aniq va tushunarli bo'lishi mumkin.
	// Satrning uzunligi aniq nazorat qilinadi.
	for i := 0; i < len(hello); i++ {
		fmt.Printf("Belgi: %c, Index:%d\n", hello[i], i)
		// fmt.Println("Birinchi belgini:", string(hello[i]), i)

	}

	// 	Oddiy va toza sintaksis.
	// UTF-8 kodlashdagi barcha belgilarni to'g'ri qaytaradi, shu jumladan ko'p baytli belgilarni (masalan, emoji yoki xalqaro belgilar).
	// Indeks va belgini to'g'ridan-to'g'ri ajratib olish imkonini beradi.
	for i, char := range hello {
		fmt.Printf("Index: %d, Belgi: %c\n", i, char)
	}
}
