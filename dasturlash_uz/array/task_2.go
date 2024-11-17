// Massiv (array) berilgan.

// O'sha massivdagi eng kichik elementni toping va konsolga chiqaring.
package main

import (
	"fmt"
)

func Maxnum(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		// fmt.Println(nums[i])
		if nums[i] < result{
			result = nums[i]
			
		}
	}

	return result
}

func main() {
	nums := []int{12, 21, 36, 71, 24, 85, 1, 30, 25} 
                         

	newLength := Maxnum(nums) 
	fmt.Println(newLength)
	
}