package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [][]int{{0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}, {0, 1, 0, 1, 0}, {0, 0, 1, 0, 0}}
	drr:=[5][5]int{}
	start:=time.Now()
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			drr[i][j]=calc(arr,i,j)
		}
	}
	end:=time.Since(start)
	fmt.Println(end)
	fmt.Println(drr)
}

func calc(arr [][]int, i int, j int) int {
	//1 alive 0 dead
	row := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	col := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	count:=0
	for k:=0;k<8;k++{
		r:=i+row[k]
		c:=j+col[k]
		if(r<0||r>=len(arr)||c<0||c>=len(arr[0])){
			continue
		}
		
			if arr[r][c]==1{
				count++
			
		}
	}
	if arr[i][j]==1{
		if count>=4{
			return 0
		}else if count==2||count==3{
			return 1
		}else{
			return 0
		}
	}else{
		if count==3{
			return 1
		}else{
			return 0
		}
	}
}
