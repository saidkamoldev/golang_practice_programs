// Topshiriq 1: Har bir elementni kvadratga ko'paytirish
// Bir slice yarating va uning har bir elementini kvadratga ko'paytiring, keyin yangi slice ni chop eting.

package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6}

	var sum int

	for _, v := range s {

		sum += v * 2

	}
	fmt.Println(sum)
}
