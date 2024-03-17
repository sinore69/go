package main

import (
	"fmt"
	"github.com/sinore69/go-packages"
	)

func main() {
	i := 0
	for ; i < 10; i++ {
		fmt.Print(i)
		if i == 5 {
			continue

		}
	}
	fmt.Println()
	fmt.Print(i)
	j:=gopackages.Print()
	fmt.Print(j)
}
