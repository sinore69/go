package main

import (
	"fmt"
	"github.com/sinore69/go-packages"
)

func new() {
	arr := []int{1, 2, 3, 4, 5}
	mutate(arr)
	fmt.Println()
	for _, j := range arr {
		fmt.Print(j)
	}
	fmt.Print(gopackages.Print())
}
func mutate(arr []int) {
	arr[0] = 100
	for _, j := range arr {
		fmt.Print(j)
	}
}
