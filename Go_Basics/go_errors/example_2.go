// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func division(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, errors.New("deivision by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := division(4, 0)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result)
// 	}
// }
