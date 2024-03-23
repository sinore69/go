package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i // Send data to the channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close the channel when done sending data
}

func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Println(n)
	}
}

func Main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go Producer(ch, &wg)
	go Consumer(ch, &wg)
	wg.Wait()
	// Wait for a keystroke before exiting
	// fmt.Println("Press any key to exit.")
	// var input string
	// fmt.Scanln(&input)
}
