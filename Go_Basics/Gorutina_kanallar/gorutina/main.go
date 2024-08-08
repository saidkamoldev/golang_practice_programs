package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	go sayHello()               // sayHello funksiyasini gorutina sifatida ishga tushadi
	time.Sleep(1 * time.Second) // Asosiy gorutina kutib turishi kerak, aks holda sayHello bajarilmay qolishi mumkun
}
