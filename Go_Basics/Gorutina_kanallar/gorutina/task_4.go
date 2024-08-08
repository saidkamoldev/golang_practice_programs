package main

import (
	"fmt"
	"sync"
	// "time"
)

func printNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done() // Gorutina tugagandan so'ng WaitGroup ni xabardor qilish
	fmt.Println("Number:", num)
}

func main() {
	var wg sync.WaitGroup // WaitGroupni yaratish

	for i := 1; i <= 3; i++ {
		wg.Add(1)              // Har bir gorutina uchun WaitGroup ga vazifa qo'shish
		go printNumber(i, &wg) // Gorutina ishga tushurish va WaitGroup ni yuborish

	}

	wg.Wait() // Barcha gorutinalar tugashini kutish
}
