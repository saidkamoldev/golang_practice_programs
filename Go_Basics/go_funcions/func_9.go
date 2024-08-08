package main

import "fmt"

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := makeMultiplier(2)
	result := double(5)
	fmt.Println("Result:", result)
}
