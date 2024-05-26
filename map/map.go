package main

import "fmt"
type check struct{
	m map[int]int
}
func main() {
	test()
}

func (c *check)try(){
	if c.m==nil{
		fmt.Println("nil")
	}else{
		fmt.Println("not nil")
	}
}

func test(){
	m:=make(map[int]int)
	m[1]=5
	m[2]=3
	m[3]=4
	for i:=0;i<len(m);i++{
		fmt.Println(m[i])
	}
}