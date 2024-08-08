package main

import "fmt"

func main() {
    // String yaratish
    var hello string = "Salom, Dunya!"

    // Stringni e'lon qilish va qiymat berish
    greeting := "Qanday ahvolda?"

    // Stringni chop etish
    fmt.Println(hello)
    fmt.Println(greeting)

    // String uzunligini topish
    fmt.Println("Hello string uzunligi:", len(hello))

    // Stringdan belgilarni olish
    fmt.Println("Birinchi belgi:", string(hello[0]))

    // Stringlarni qo'shish (concatenation)
    combined := hello + " " + greeting
    fmt.Println("Qo'shilgan string:", combined)

    // Stringda iteratsiya qilish
    for i, char := range hello {
        fmt.Printf("Index: %d, Belgi: %c\n", i, char)
    }
}
