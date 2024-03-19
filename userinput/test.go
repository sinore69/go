package main

import (
	"fmt"
	"os"
)

func main() {
	args:=os.Args[1:]
	arr:=append(args, "new")
	fmt.Print(arr[3])
}
