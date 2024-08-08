package main

import "fmt"

// Person nomli struct ni e'lon qilamiz
type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// struct ning bo'sh qiymatlarini yaratish
	var p1 Person
	fmt.Println(p1) // { 0 }

	// struct ga qiymat berish
	p1.Name = "Alice"
	p1.Age = 30
	p1.Email = "alice@example.com"
	fmt.Println(p1) // {Alice 30 alice@example.com}

	// struct ni to'ldirib yaratish
	p2 := Person{
		Name:  "Bob",
		Age:   25,
		Email: "bob@example.com",
	}
	fmt.Println(p2) // {Bob 25 bob@example.com}

	// struct ni qisqa usulda yaratish
	p3 := Person{"Charlie", 35, "charlie@example.com"}
	fmt.Println(p3) // {Charlie 35 charlie@example.com}

	// struct ning maydonlarini o'qish
	fmt.Println("Name:", p3.Name)
	fmt.Println("Age:", p3.Age)
	fmt.Println("Email:", p3.Email)
}
