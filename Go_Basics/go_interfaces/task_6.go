package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Human struct {
	name string
}

func (h Human) Speak() string {
	return "Hello, my name is" + h.name
}

func main() {
	var s1, s2 Speaker

	s1 = Human{name: "Alice"}
	s2 = Human{name: "Alice"}

	fmt.Println(s1 == s2) //true

	s2 = Human{name: "Bob"}

	fmt.Println(s1 == s2) // False
}
