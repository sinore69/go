package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i // Send data to the channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close the channel when done sending data
}

func consumer(ch <-chan int) {
    val, isopen :=<-ch
    for isopen{
            fmt.Println(val)
            val, isopen=<-ch
        }
}

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	// Wait for a keystroke before exiting
	fmt.Println("Press any key to exit.")
	var input string
	fmt.Scanln(&input)
}
