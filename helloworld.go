package main

import (
	"fmt"
	// "github.com/sinore69/go-packages"
	// "github.com/sinore69/go/concurrency"
	"github.com/sinore69/go/quicksort"
)

// func new() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	mutate(arr)
// 	fmt.Println()
// 	for _, j := range arr {
// 		fmt.Print(j)
// 	}
// 	fmt.Print(gopackages.Print())
// }
// func mutate(arr []int) {
// 	arr[0] = 100
// 	for _, j := range arr {
// 		fmt.Print(j)
// 	}
// }

func main(){
	// concurrency.Main();
	// new()
	arr:=[]int{54,7,9,3,23,7,9,4,2,6,9,84,4}
	newarr:=quicksort.Quicksort(arr)
	fmt.Println(newarr)
}