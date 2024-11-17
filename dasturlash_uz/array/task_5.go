// n, array[] berilgan.

// Shu massivda n-chi index dagi elementini konsolga chiqaring.

// Agar index xato bo'lsa  0 chiqaring.
package main

import (
	"fmt"
)

func Maxnum(nums []int, target int) int {
	
	for i := 0; i < len(nums); i++ {
		if target == i {
			return nums[i]
		}
	}

	return 0
}

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6}                
	target := 3
	newLength := Maxnum(nums, target) 
	fmt.Println(newLength)
	
}