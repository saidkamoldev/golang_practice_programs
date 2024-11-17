// n, array[] berilgan. 

// Shu massivda n soni bor yoki  yo'qligini aniqlang. n soni bo'lsa true, bo'lmasa false konsolga chiqaring.

package main

import (
	"fmt"
)

func Maxnum(nums []int, target int) int {
	
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return 1
		}
	}

	return 0
}

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6}                
	target := 28
	newLength := Maxnum(nums, target) 
	fmt.Println(newLength)
	
}