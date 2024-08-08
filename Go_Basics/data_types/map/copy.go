package main

import (
	"fmt"
	// "strings"
)

func main() {
	// mapni elon qilish va yaratish
	var m map[string]int
	m = make(map[string]int)

	// Qisqa usulda mapni elon qilish
	n := make(map[string]int)

	// Mapni to'ldirish
	m["Alice"] = 25
	m["Bob"] = 30

	n["Charlie"] = 35
	n["Dave"] = 40

	//Map dagi qiymatlarni o'qish
	fmt.Println("Alice's age:", m["Alice"])
	fmt.Println("Bob's age:", m["Bob"])

	// Mapdagi qiymatni yangilash
	m["Alice"] = 25
	fmt.Println("Alice's new age:", m["Alice"])

	//Mapdagi elementlarni iteratsiya qilish
	for key, value := range n {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
