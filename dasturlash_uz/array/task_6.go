// n, array[] berilgan.

// Shu massiv oxiriga n ni jo'ylashtiring va massivni konsolga chiqaring.
package main

import (
	"fmt"
)

func Maxnum(nums []int, target int) []int {
	result := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		result[i]= nums[i]
	}

	result[len(nums)]=target
	return result
}

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6}    
	target:= 13            
	newLength := Maxnum(nums, target) 
	fmt.Println(newLength)
	
}