// Topshiriq 4: Elementlarni teskari tartibda chiqarish
// Bir slice yarating va uning elementlarini teskari tartibda chop eting.

package main

import "fmt"

func main() {
	// Slice yaratish
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Teskari tartibdagi slice ni yaratish uchun bo'sh slice
	var reversed []int

	// Slice ustida teskari iteratsiya qilish
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}

	// Teskari tartibdagi slice ni chop etish
	fmt.Println("Teskari tartibdagi slice:", reversed)
}
