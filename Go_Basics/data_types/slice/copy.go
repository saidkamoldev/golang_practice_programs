package main

import "fmt"

func main() {
	// Bo'sh slice yaratish
	var s []int
	fmt.Println("Bo'sh slice:", s)

	// Elementlari bilan slice yaratish
	s = []int{1, 2, 3, 4, 5}
	fmt.Println("Elementlari bilan slice:", s)

	// Qisqa usulda slice yaratish
	t := []string{"hello", "world"}
	fmt.Println("Qisqa usulda slice:", t)

	// Slice dan qism olish
	u := s[1:4] // 2, 3, 4 elementlarini oladi
	fmt.Println("Slice dan qism:", u)

	// Slice ni uzunligini va sig'imi
	fmt.Println("Uzunligi:", len(s))
	fmt.Println("Sig'imi:", cap(s))

	// Slice ga element qo'shish
	s = append(s, 6)
	fmt.Println("Element qo'shishdan keyin slice:", s)

	// Slice ustida iteratsiya qilish
	for i, v := range s {
		fmt.Printf("Index: %d, Qiymat: %d\n", i, v)
	}
}
