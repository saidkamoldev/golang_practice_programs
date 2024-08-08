package main

import (
	"fmt"
	"time"
)

func printNumber(num int) {
	fmt.Println("Number:", num)
}

func main() {
	go printNumber(99) // Parametr bilan funcksiyani gorutina sifatida ishga tushurish
	time.Sleep(1 * time.Second)
}
