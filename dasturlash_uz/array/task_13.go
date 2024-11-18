// int array1[], int array2[] berilgan.

// Shu ikkita arraydagi  o'xshash elementlarni ekranga chiqaring.
package main

import (
	"fmt"
	"strconv"
)

func TwoSum(nums1 []int, nums2 []int) []int {

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				fmt.Println("result[" + strconv.Itoa(i) + "]: " + strconv.Itoa(nums2[j]) + " = result[" + strconv.Itoa(j) + "]: " + strconv.Itoa(nums2[j]))
			}
		}
	}


    // result := []int{}
	// index := 0
	// for i := 0; i < len(nums); i++ {
	// 	for j := i +1; j < len(nums); j++ {
	// 		if nums[j] == nums[i] {
    //             fmt.Println("result[" + strconv.Itoa(i) + "]: " + strconv.Itoa(nums[i]) + " = result[" + strconv.Itoa(j) + "]: " + strconv.Itoa(nums[j]))
	// 		}
	// 	}
	// }
	



	return nums1
}

func main() {
	nums1 := []int{1,2,3}    
	nums2 := []int{1,4,1,2}   

	newLength := TwoSum(nums1, nums2) 
	
	fmt.Println(newLength)
	
}