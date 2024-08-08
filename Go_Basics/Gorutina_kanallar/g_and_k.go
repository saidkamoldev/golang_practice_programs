package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()     // Gorutina tugagandan so'ng wg.Done() chaqiriladi
	for n := range ch { // Kanal yopilmaguncha ma'lumotni o'qish
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func main() {
	const numWorkers = 3
	ch := make(chan int)  // Int turidagi kanal yaratish
	var wg sync.WaitGroup // WaitGroup yaratish

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)             // Har bir worker uchun WaitGroup ga vazifa qo'shish
		go worker(i, ch, &wg) // Worker funksiyasini gorutina sifatida ishga tushirish
	}

	for j := 0; j < 5; j++ {
		ch <- j // Kanaldan workerlarga ma'lumot yuborish
	}
	close(ch) // Kanalni yopish, bu workerlarga endi yangi ma'lumot kelmasligini bildiradi
	wg.Wait() // Barcha workerlar tugashini kutish
}
