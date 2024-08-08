// Topshiriq 5: Berilgan qiymatlarni o'chirish
// Bir slice yarating va undan ma'lum bir qiymatni o'chiring, keyin yangi slice ni chop eting.

package main

import "fmt"

func main() {
	// Boshlang'ich slice yaratish
	s := []int{1, 2, 3, 4, 5, 3, 6, 3}

	// O'chirilishi kerak bo'lgan qiymat
	valueToRemove := 3

	// Natijaviy slice uchun bo'sh slice yaratish
	var result []int

	// Boshlang'ich slice ustida iteratsiya qilish
	for _, v := range s {
		if v != valueToRemove {
			result = append(result, v)
		}
	}

	// Natijaviy slice ni chop etish
	fmt.Println("Qiymat o'chirilgandan keyin slice:", result)
}
