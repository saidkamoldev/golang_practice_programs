package main

import "fmt"

func main() {
	// map ni e'lon qilish va yaratish
	var m map[string]int
	m = make(map[string]int)

	// Qisqa usulda map ni e'lon qilish
	n := make(map[string]int)

	// Map ni to'ldirish
	m["Alice"] = 25
	m["Bob"] = 30

	n["Charlie"] = 35
	n["Dave"] = 40

	// Map dagi qiymatlarni o'qish
	fmt.Println("Alice's age:", m["Alice"])
	fmt.Println("Bob's age:", m["Bob"])

	// Mapdagi qiymatni yangilash
	m["Alice"] = 26
	fmt.Println("Alice's new age:", m["Alice"])

	// Mapdan qiymatni o'chirish
	delete(m, "Bob")
	fmt.Println("Bob's age after deletion:", m["Bob"])

	// Mapdagi elementlarni iteratsiya qilish
	for key, value := range n {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
