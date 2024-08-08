package main

import (
	"fmt"
	"sync"
)

// Birinchi gorutina funsksiyasi

func workerA(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Gorutina tugaganidan so'ng WaitGroupni ogohlantiramiz
	for n := range ch {
		fmt.Printf("Worker A %d received %d and processed it to %d\n", id, n, n*2)
	}
}

// Ikkinchi Gorutina funksiyasi
func workerB(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Gorutina tugaganidan so'ng WaitGroupni ogohlantirish
	for n := range ch {
		fmt.Printf("Worker B %d received %d and processed it to %d\n", id, n, n*n)
	}

}

func main() {
	const numWorkersA = 2
	const numWorkersB = 2
	ch := make(chan int)
	var wg sync.WaitGroup

	// Worker A gorutinalarni ishga tushirish
	for i := 0; i < numWorkersA; i++ {
		wg.Add(1)
		go workerA(i, ch, &wg)
	}

	// Worker b gorutinalarni ishga tushirish
	for i := 0; i < numWorkersB; i++ {
		wg.Add(1)
		go workerB(i, ch, &wg)
	}

	for j := 1; j <= 10; j++ {
		ch <- j
	}

	close(ch) // Kanalni yopish
	wg.Wait() // Barcha gorutinalarni tugatish
}
