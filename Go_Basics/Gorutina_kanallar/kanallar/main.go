package main

import (
	"fmt"
)

func sendMessage(ch chan string) {
	ch <- "Hello from goroutine" // Kanalga xabar yuborish
}

func main() {
	ch := make(chan string)

	go sendMessage(ch) // sendMessage funsksiyasini gorutina sifatida ishga tushirish

	msg := <-ch // Kanaldan xabar o'qish
	fmt.Println(msg)
}
