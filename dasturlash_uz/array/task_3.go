// Massive (array) berilgan.

// Massive elementlarini o'rtacha qiymatlarini hisoblang va konsolga chiqaring.
package main

import (
	"fmt"
)

func Maxnum(nums []int) int {
	result := 1
	for i := 0; i < len(nums); i++ {
		// fmt.Println(nums[i])
		result += nums[i]
	}
	sum := result / len(nums)
	return sum
}

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6} 
                         

	newLength := Maxnum(nums) 
	fmt.Println(newLength)
	
}