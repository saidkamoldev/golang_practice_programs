// Massiv (array) berilgan.

// O'sha massivdagi eng katta elementni toping va konsolga chiqaring.
package main

import (
	"fmt"
)

func Maxnum(nums []int) int {
	result := nums[0]
	for i := 0; i < len(nums); i++ {
		// fmt.Println(nums[i])
		if nums[i] > result{
			result = nums[i]
			
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, 2} 
                         

	newLength := Maxnum(nums) 
	fmt.Println(newLength)
	
}