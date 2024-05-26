package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done:=make(chan bool)
	go read(ch)
	go write1(ch,done)
	<-done
	//go write2(ch)
}

func write1(ch chan<- int,done chan<-bool) {
	defer close(ch)
	time.Sleep(time.Millisecond*000)
	fmt.Println("sending data from write1")
	ch <- 69
	done<-true
}

// func write2(ch chan<- int) {
// 	defer close(ch)
// 	time.Sleep(time.Millisecond*4000)
// 	fmt.Println("sending data from write2")
// 	ch <- 420
// }

func read(ch <-chan int){
	fmt.Println("channel listening")
	for num:= range ch{
		fmt.Println(num)
	}
}