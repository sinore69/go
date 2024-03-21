package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	start := time.Now()

	go delay1(&wg)
	go delay2(&wg)
	delay3(&wg)
	wg.Wait()
	elapsed := time.Now()
	fmt.Println(elapsed.Sub(start))
}
func delay1(wg *sync.WaitGroup) {
	defer wg.Done()
	// start := time.Now()
	for i := 0; i < 100000000; i++ {
	}
	// elapsed := time.Now()
	// fmt.Println(elapsed.Sub(start))
}
func delay2(wg *sync.WaitGroup) {
	defer wg.Done()
	// start := time.Now()
	for i := 0; i < 100000000; i++ {
	}
	// elapsed := time.Now()
	// fmt.Println(elapsed.Sub(start))
}
func delay3(wg *sync.WaitGroup) {
	defer wg.Done()
	// start := time.Now()
	for i := 0; i < 100000000; i++ {
	}
	// elapsed := time.Now()
	// fmt.Println(elapsed.Sub(start))
}
