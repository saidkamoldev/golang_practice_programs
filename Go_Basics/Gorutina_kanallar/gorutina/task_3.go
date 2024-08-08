package main

import (
	"fmt"
	"time"
)

func printNumber(num int) {
	fmt.Println("Number:", num)
}

func main() {
	for i := 1; i <= 5; i++ {
		go printNumber(i) // Har bir iteratsiyada yangi gorutina ishga tushirish

	}
	time.Sleep(1 * time.Second) //Asosiy gorutina tugamasligi uchun kutish
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printNumber(num int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Number:", num)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go printNumber(i, &wg)
// 	}

// 	wg.Wait() // Barcha gorutinalar tugashini kutish
// }
