package main

import "fmt"

func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}
func main() {
	result := apply(add, 3, 4)
	fmt.Println("Result:", result)
}
