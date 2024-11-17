package main

import (
	"fmt"
)

func RemoveElement(nums []int, target int) []int {
	result := make([]int, len(nums)-1)
	index := 0
	for i := 0; i < len(nums); i++ {
		if i != target {
			result[index] = nums[i]
			index++
		}

	}


	return result
}

func main() {
	nums := []int{4,5,6,7}    
	target:= 2         
	newLength := RemoveElement(nums, target) 
	fmt.Println(newLength)
	
}