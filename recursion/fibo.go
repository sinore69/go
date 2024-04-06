package main

import "fmt"

func main() {
	n := 500
	m := make(map[int]int)
	res := fibonacci(n, m)
	fmt.Printf("%v", res)
}

func fibonacci(n int, m map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 2 || n == 1 {
		return 1
	}
	if val, present := m[n]; present {
		return val
	} else {
		m[n] = fibonacci(n-2, m) + fibonacci(n-1, m)
	}
	return fibonacci(n-1, m) + fibonacci(n-2, m)
}
