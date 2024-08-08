package main

import (
	"fmt"
	"time"
)

func printNumber(ch chan int) {
	num := <-ch //Kanaldan sonni qabul qilish
	fmt.Println("Recevied number:", num)
}

func printNumber2(son chan int) {
	raqam := <-son
	fmt.Println("2 Recevied number:", raqam)
}

func main() {
	ch := make(chan int)  // Kanal yaratish
	son := make(chan int) // Ikkinchi kanalni yaratish
	go printNumber(ch)    // Gorutina ishga tushirish
	go printNumber2(son)  // 2 chi Gorutinani ishga tushurish
	ch <- 42              // Kanalga sonni yuborish
	son <- 100
	time.Sleep(1 * time.Second)
}
