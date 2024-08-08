// Vazifa 1: Sonlarni kvadratga ko'paytirish
// Yangi squareWorker funksiyasini yarating, u kanaldan sonlarni olib, ularni kvadratga ko'paytirib, natijani boshqa kanalda chop etsin.

// Talablar:

// 3 ta squareWorker gorutinalarini yarating.
// Bir kanal orqali sonlarni squareWorker gorutinalariga yuboring.
// squareWorker gorutinalar kvadratga ko'paytirilgan sonlarni boshqa kanalga yuborishi kerak.
// Asosiy gorutina kvadratga ko'paytirilgan sonlarni chop etsin.

package main

import (
	"fmt"
	"sync"
)

func squareWorker(id int, input chan int, output chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range input {
		result := n * n
		fmt.Printf("Worker %d processed %d to %d\n", id, n, result)
		output <- result

	}

}

func main() {
	const numWorkers = 3
	input := make(chan int)
	output := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go squareWorker(i, input, output, &wg)
	}

	go func() {
		for j := 1; j <= 5; j++ {
			input <- j
		}

		close(input)
	}()

	go func() {
		wg.Wait()
		close(output)
	}()

	for result := range output {
		fmt.Println("Result:", result)
	}
}
