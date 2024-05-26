package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	left, right := 0, len(arr)-1
	ele := 69
	found:=false
	var mid int
	for left<=right{
		mid = left + (right-left)/2
		if arr[mid] == ele {
			fmt.Printf("element found on index %d",mid)
			found=true
			break
		}
		if ele<arr[mid]{
			right=mid-1
		}
		if ele>arr[mid]{
			left=mid+1
		}

	}
	if !found{
		fmt.Println("not found")
	}

}