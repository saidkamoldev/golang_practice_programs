// int array1[] berilgan.

// Shu arraydagi o'xshash elementlarni olib tashang.
package main

import (
	"fmt"
	// "strconv"
)

func FindCommonElements(nums1 *[]int) int {

	nums := *nums1
	temp := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[temp] {
			nums[temp+1] = nums[i]
			temp++
		}
	}


	return temp+1 
}

func main() {
	nums := []int{1, 1, 2, 2, 5, 3, 6, 6, 7, 7, 7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,}

    len :=	FindCommonElements(&nums)

	for i := 0; i < len; i++ {
		fmt.Print(nums[i])
	}

}
