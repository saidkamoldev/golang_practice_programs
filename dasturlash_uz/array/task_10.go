// array[] berilgan.

// Shu massivni teskarisini yangi massivga ko'chirib o'tkazing va konsolga chiqaring.
package main

import (
	"fmt"
)

func ReversevArray(nums []int) []int {
	result := make([]int, len(nums))
	index := 0
	for i := len(nums)-1; i >= 0; i-- {
		result[index] = nums[i]
		index++
	}


	return result
}

func main() {
	nums := []int{4,5,6,7}    

	newLength := ReversevArray(nums) 
	
	fmt.Println(newLength)
	
}