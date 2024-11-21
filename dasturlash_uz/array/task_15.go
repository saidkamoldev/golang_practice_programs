// Shu massivni  elementlarini o'sish tartibida jo'ylashtiring (sort by increasing (ascending) ) va hosil bo'lgan massivni ni konsolga chiqaring.
package main

import (
	"fmt"
)

func SortAscending(nums []int) []int {
	// array := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[j]
				nums[j] = nums[i]
				nums[i] = temp
			}
		}

	}

	return nums
}

func main() {
	nums := []int{2, 4, 1, 5}

	newLength := SortAscending(nums)

	fmt.Println(newLength)

}
