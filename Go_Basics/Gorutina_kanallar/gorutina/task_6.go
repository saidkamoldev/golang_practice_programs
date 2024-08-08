package main

import (
	"fmt"
	"sync"
)

func printNumber(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Gorutina tugaganidan so'ng WaitGroup ni xabardor qilish
	num := <-ch     // Kanaldan malumotlarni qabul qilish
	fmt.Println("Received number:", num)
}

func printNumber2(num chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	son := <-num
	fmt.Println("Qabul qilingan raqam:", son)
}

func main() {
	var wg sync.WaitGroup // WaitGroup yaratish
	ch := make(chan int)  // Kanalni yaratish
	num := make(chan int) // 2 kanal yartish
	for i := 1; i <= 3; i++ {
		wg.Add(1)               // Har bir gorutina uchun WaitGroup ga vazifa qushish
		go printNumber(ch, &wg) // Gorutinani ishga tushurish va WaitGroup ni yuborish

	}
	for j := 1; j <= 3; j++ {
		ch <- j // Kanallarga sonlarni yuborish
	}

	for s := 1; s <= 5; s++ {
		wg.Add(1)
		go printNumber2(num, &wg)
		num <- s
	}

	close(ch) // Kanalni yopish
	wg.Wait() // Barcha gorutinalarni tugatishini kutish
}

// package main

// import (
//     "fmt"
//     "sync"
// )

// // Birinchi gorutina funksiyasi
// func workerA(id int, ch chan int, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for n := range ch {
//         fmt.Printf("Worker A %d received %d and processed it to %d\n", id, n, n*2)
//     }
// }

// // Ikkinchi gorutina funksiyasi
// func workerB(id int, ch chan int, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for n := range ch {
//         fmt.Printf("Worker B %d received %d and processed it to %d\n", id, n, n*n)
//     }
// }

// func main() {
//     const numWorkersA = 2
//     const numWorkersB = 2
//     ch := make(chan int)
//     var wg sync.WaitGroup

//     // Worker A gorutinalarni ishga tushirish
//     for i := 0; i < numWorkersA; i++ {
//         wg.Add(1)
//         go workerA(i, ch, &wg)
//     }

//     // Worker B gorutinalarni ishga tushirish
//     for i := 0; i < numWorkersB; i++ {
//         wg.Add(1)
//         go workerB(i, ch, &wg)
//     }

//     // Kanaldan sonlarni yuborish
//     for j := 1; j <= 10; j++ {
//         ch <- j
//     }

//     close(ch) // Kanalni yopish
//     wg.Wait() // Barcha gorutinalar tugashini kutish
// }
