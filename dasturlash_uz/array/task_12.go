// array[] berilgan.

// Shu arraydagi o'xshash elementlarni ekranga chiqaring.
package main

import (
	"fmt"
	"strconv"
)

func SimilarElement(nums []int) []int {
    // result := []int{}
	// index := 0
	for i := 0; i < len(nums); i++ {
		for j := i +1; j < len(nums); j++ {
			if nums[j] == nums[i] {
                fmt.Println("result[" + strconv.Itoa(i) + "]: " + strconv.Itoa(nums[i]) + " = result[" + strconv.Itoa(j) + "]: " + strconv.Itoa(nums[j]))
			}
		}
	}
	


	return nums
}

func main() {
	nums := []int{1,2,1,4}    

	newLength := SimilarElement(nums) 
	
	fmt.Println(newLength)
	
}