// Topshiriq 2: Juft va toq sonlarni ajratish
// Bir slice yarating va uni juft va toq sonlarga ajrating, har bir yangi slice ni alohida chop eting.

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var juft []int
	var toq []int

	for _, v := range s {
		if v%2 == 0 {
			juft = append(juft, v)
		} else {
			toq = append(toq, v)
		}
	}

	fmt.Println("Juft sonlar:", juft)
	fmt.Println("Toq sonlar:", toq)
}
