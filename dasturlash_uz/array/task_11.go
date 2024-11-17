// n va array[] berilgan.

// Shu massivda nechta n soni borligini topib konsolga chiqaring.

// Masalan:  n = 4 , array[2,4,6,4,] 
package main

import (
	"fmt"
)

func ElementCount(nums []int, target int) int {

	count :=0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}

	return count
}

func main() {
	nums := []int{2,4,6,4}    
	target := 4
	newLength := ElementCount(nums, target) 
	
	fmt.Println(newLength)
	
}