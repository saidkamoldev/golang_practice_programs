// a, array berilgan.

// Arrayning xohlagan ikkita  elementlarini  yig'indisi a ga teng bo'lishini aniqlang va ularni  konsolga chiqaring.
// GetPairForSum
package main

import (
	"fmt"
)

func GetPairForSum(nums []int, target int) []int {
	// array := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j]  == target{
			 fmt.Println(nums[i], " + ", nums[j], " = ", target)
			}
		}

	}

	return nums
	// for i := 0; i < len(nums); i++ {

	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] < nums[j] {
	// 			temp := nums[j]
	// 			nums[j] = nums[i]
	// 			nums[i] = temp
	// 		}
	// 	}

	// }

}

func main() {
	nums := []int{1,2,3,7,4}
	target := 5
	newLength := GetPairForSum(nums, target)

	fmt.Println(newLength)

}