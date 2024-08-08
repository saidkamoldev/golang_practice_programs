package main

import "fmt"

// func main() {
// 	// // struct ning bo'sh qiymatlarini yaratish
// 	// var p1 Person
// 	// // fmt.Println(p1) // { 0 }

// 	// // struct ga qiymat berish
// 	// // p1.Name = "Alice"
// 	// // p1.Age = 30
// 	// // p1.Email = "test@gmail.com"
// 	// // fmt.Println(p1) //{Alice 30 test@gmail.com}
// 	// for i := 0; i <= 10; i++ {
// 	// 	fmt.Println("1.yangi foydalanuvchi qushish\n 2.oldingi foydalanuvchilarni ko'rish\n 3 .exit")
// 	// }

// 	// fmt.Println("Ismingiz, yoshingiz va emailingizni kiriting (bir qatorni bo'sh joy bilan ajratib):")
// 	// _, err := fmt.Scan(&p1.Name, &p1.Age, &p1.Email)
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }

// 	// fmt.Println(p1)

// }

// Person nomli struct ni elon qilib olamiz
type Person struct {
	Name  string
	Age   int
	Email string
}

func clearPerson(p *Person) {
	p.Name = ""
	p.Age = 0
	p.Email = ""
}

func users() {
	var people []Person
	condition := true
	var sum int

	for condition {
		fmt.Println("1. Yangi foydalanuvchi qo'shish\n2. Oldingi foydalanuvchilarni ko'rish\n3. Chiqish\n4. Foydalanuvchini tozalash")
		fmt.Scan(&sum)

		if sum == 3 {
			condition = false
		} else if sum == 1 {
			var p Person
			fmt.Printf("Ism, yosh va emailni kiriting:\n")
			fmt.Scan(&p.Name, &p.Age, &p.Email)
			people = append(people, p)
		} else if sum == 2 {
			fmt.Println("Kiritilgan foydalanuvchilar:")
			for _, p := range people {
				fmt.Printf("Ism: %s, Yosh: %d, Email: %s\n", p.Name, p.Age, p.Email)
			}

		} else if sum == 4 {
			if len(people) > 0 {
				clearPerson(&people[0])
				fmt.Println("Birinchi foydalanuvchi tozalandi.")
			} else {
				fmt.Println("Noto'g'ri tanlov, qayta urinib ko'ring.")
			}
		}
	}
}

func main() {
	users()
}

// Albatta! Ushbu kodda append va range funksiyalarining roli quyidagicha:

// append funksiyasi:
// append funksiyasi Go tilida slice'ga (kesma) yangi element qo'shish uchun ishlatiladi. people slice'siga yangi Person tuzilmasini qo'shish uchun foydalaniladi.

// go
// Копировать код
// var p Person
// fmt.Printf("Ism, yosh va emailni kiriting:\n")
// fmt.Scan(&p.Name, &p.Age, &p.Email)
// people = append(people, p)
// Ushbu kod parchasida:

// var p Person - yangi Person ob'ekti p yaratiladi.
// fmt.Printf("Ism, yosh va emailni kiriting:\n") - foydalanuvchidan ism, yosh va emailni kiritishni so'raydi.
// fmt.Scan(&p.Name, &p.Age, &p.Email) - foydalanuvchining kiritgan ma'lumotlarini p ob'ektining Name, Age, va Email maydonlariga saqlaydi.
// people = append(people, p) - people slice'siga yangi p ob'ektini qo'shadi.
// append funksiyasi people slice'sini kengaytiradi va oxiriga yangi elementni qo'shadi. appendning ishlashi:

// Agar slice'da yetarli joy bo'lsa, yangi element qo'shiladi va slice uzunligi oshadi.
// Agar yetarli joy bo'lmasa, yangi slice yaratilib, eski slice elementlari va yangi element unga nusxalanadi.
// range kalit so'zi:
// range kalit so'zi slice yoki map kabi kolleksiyalarni iteratsiya qilish uchun ishlatiladi. people slice'sidagi barcha elementlarni iteratsiya qilish uchun foydalaniladi.

// go
// Копировать код
// fmt.Println("Kiritilgan foydalanuvchilar:")
// for _, p := range people {
//     fmt.Printf("Ism: %s, Yosh: %d, Email: %s\n", p.Name, p.Age, p.Email)
// }
// Ushbu kod parchasida:

// for _, p := range people - people slice'sidagi har bir element ustidan iteratsiya qilinadi. Har bir iteratsiyada p o'zgaruvchisi people slice'sidagi navbatdagi Person ob'ektini oladi.
// fmt.Printf("Ism: %s, Yosh: %d, Email: %s\n", p.Name, p.Age, p.Email) - har bir Person ob'ekti uchun ism, yosh va emailni chop etadi.
// range ishlashi:

// range people har bir iteratsiyada ikki qiymatni qaytaradi: indeks va element. Bu yerda biz indeksni ishlatmayapmiz, shuning uchun _ bilan uni bekor qildik. p esa har bir elementni oladi.
// Umuman, append va range yordamida biz slice'ga yangi elementlarni qo'shishimiz va slice ichidagi barcha elementlarni iteratsiya qilib ularga ishlov berishimiz mumkin.
