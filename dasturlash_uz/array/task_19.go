// int a[], int b[] massivlari berilgan.

// Shu massivlarni o'zaro teng ekanligini aniqlang. Agar ular teng bo'lsa true . Bo'lmasa false konsolga chiqaring.

// O'zaro teng deganda ularning barcha elementlari tengligi nazarda tutilyabdi.  Arraylarning uzunligi bir xil deb oling.
// AreArraysEqual

package main

import (
	"fmt"
)

func AreArraysEqual(array_a []int, array_b []int) bool{
	// array := make([]int, len(nums))
	var status bool
	for i := 0; i < len(array_a); i++ {
		if array_a[i] == array_b[i] {
			status = true
		} else {
			status = false
			break
		}
	}
	return status
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
	array_a := []int{1,2,3,7,4}
	array_b :=[]int{1,2,4,7,4}

	newLength := AreArraysEqual(array_a, array_b)

	fmt.Println(newLength)

}