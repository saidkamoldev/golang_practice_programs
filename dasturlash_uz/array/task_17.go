// GetSecondMax
// int array[] berilgan.

// O'sha massivdagi eng katta ikkinchi  elementni toping. Bunda array elementlari bir biriga teng emas deb oling.
package main

import (
	"fmt"
)

func GetSecondMax(nums []int) int {
	// array := make([]int, len(nums))
	max := nums[0]
	answer := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			answer = max
			max = nums[i]
		}
	}
	// for i := 0; i < len(nums); i++ {

	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] < nums[j] {
	// 			temp := nums[j]
	// 			nums[j] = nums[i]
	// 			nums[i] = temp
	// 		}
	// 	}

	// }

	return answer
}

func main() {
	nums := []int{1,2,6,4,7,5}

	newLength := GetSecondMax(nums)

	fmt.Println(newLength)

}