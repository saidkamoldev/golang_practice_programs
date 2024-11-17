// n, array[] berilgan.

// Shu massiv boshiga n ni jo'ylashtiring va konsolga chiqaring

package main

import (
	"fmt"
)

func Maxnum(nums []int, target int) []int {
	result := make([]int, len(nums)+1)
	 result[0] = target
	 size := 0
	for i := 1; i < len(nums)+1; i++ {
		// result[0] = target

		result[i] = nums[size]

		size++
	}


	// result[len(nums)]=target
	return result
}

func main() {
	nums := []int{6, 25, 91, 23, 72, 9, 18, 6}    
	target:= 13            
	newLength := Maxnum(nums, target) 
	fmt.Println(newLength)
	
}