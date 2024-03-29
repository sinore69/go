package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var head, tail, def *Node

func (N *Node) insert(data int) {
	newnode := Node{data: data, next: def}
	if head == tail && tail == def {
		head = &newnode
		tail = &newnode
	} else {
		tail.next = &newnode
		tail = tail.next
	}
}

func main() {
	var linkedlist = Node{}
	linkedlist.insert(1)
	linkedlist.insert(2)
	linkedlist.insert(3)

	temp := head
	for temp != def {
		fmt.Println(temp.data)
		temp = temp.next
	}

}
