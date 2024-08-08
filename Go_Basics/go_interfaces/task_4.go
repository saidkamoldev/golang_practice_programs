package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
	describe(42)
	describe("hello")
	describe(true)
}
