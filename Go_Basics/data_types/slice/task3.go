// Topshiriq 3: Eng katta va eng kichik qiymatni topish
// Bir slice yarating va uning ichida eng katta va eng kichik qiymatni toping va chop eting.

package main

import "fmt"

func main() {
	// Slice yaratish
	s := []int{12, 3, 5, 7, 19, 1, 10}

	// Dastlabki eng katta va eng kichik qiymatlarni slice ning birinchi elementi bilan e'lon qilish
	max := s[0]
	min := s[0]

	// Slice ustida iteratsiya qilish
	for _, v := range s {
		if v > max { // Agar jory element max dan katta bo'lsa
			max = v // max ni jory element bilan yangilash
		}
		if v < min { // Agar jory element min dan kichik bo'lsa
			min = v // min ni jory element bilan yangilash
		}
	}

	// Natijalarni chop etish
	fmt.Printf("Eng katta qiymat: %d\n", max)
	fmt.Printf("Eng kichik qiymat: %d\n", min)
}
