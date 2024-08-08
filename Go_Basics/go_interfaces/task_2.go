package main

import "fmt"

// Speaker interfeysi
type Speaker interface {
	Speak() string
	SpeakAge() int
}

// Human struct
type Human struct {
	name string
	age  int
}

func (h Human) Speak() string {
	return "Hello, my name is" + h.name
}

func (h Human) SpeakAge() int {
	return h.age
}

// Dog struct
type Dog struct {
	name string
	age  int
}

func (d Dog) Speak() string {
	return "Woof! My name is" + d.name
}

func (d Dog) SpeakAge() int {
	return d.age
}

func main() {
	var speakers []Speaker
	speakers = append(speakers, Human{name: "Alice", age: 21})
	speakers = append(speakers, Dog{name: "Buddy", age: 23})

	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
		fmt.Printf("I am %d years old\n", speaker.SpeakAge())
	}
}
