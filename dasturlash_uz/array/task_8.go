// a,b, array[] berilgan.

// Shu massivning a indexsiga b sonini joylashtiring  va massivni konsolga chiqaring.
package main

import (
	"fmt"
)

func Maxnum(nums []int, target int, add int) []int {
	result := make([]int, len(nums)+1)
	 size := len(result)
	index := 0
	for i := 0; i < size; i++ {
		if i == target {
			result[i] = add
		} else {
			result[i] = nums[index]
			index++
		}
	}
	// for i := 1; i < len(nums)+1; i++ {
	// 	if i == target {
	// 		result[i] = add
	// 	} else {
	// 		result[i] = nums[size]
	// 	}
		
	// 	size++
	// }
	// for i := 0; i < len(result); i++ {
	// 	if result[i] == 0 {
			
	// 	}
	// }
	// result[0] = nums[0]
	// for i := 1; i < len(nums)+1; i++ {
	// 	if i == target{
	// 		result[size] = add
	// 	}else {
	// 		result[size] = nums[size]
	// 		size++
	// 	}
		
	// }


	return result
}

func main() {
	nums := []int{4,5,6,7}    
	target:= 2 
	add := 9           
	newLength := Maxnum(nums, target, add) 
	fmt.Println(newLength)
	
}