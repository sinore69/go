package main

import (
	"fmt"
	"time"
)
type foo struct{
	name string
}
func main() {
	// args:=os.Args[1:]
	// arr:=append(args, "new")
	// fmt.Print(arr[3])
	data:=foo{}
	test(&data)
}
func test(data *foo){
	if data==nil{
		fmt.Println(data)
	}
}
func bench(){
	start:=time.Now()
	arr:=[116][116]int{}

	//for k:=0;k<60;k++{
	for i:=0;i<115;i++{
		for j:=0;j<115;j++{
			arr[i][j]=1
		}
	}
	//time.Sleep(10*time.Millisecond)
	for i:=0;i<15;i++{
		for j:=0;j<15;j++{
			arr[i][j]=0
		}
	//}
}
	elapsed:=time.Since(start)
	fmt.Println(arr)
	fmt.Println(elapsed)
}