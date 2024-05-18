package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
	fmt.Println(arr)
}

func heapify(arr []int, n int, i int) {
	left := 2 * i
	right := 2*i + 1
	largest := i
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, largest, i)
		heapify(arr,n,largest)
	}
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
