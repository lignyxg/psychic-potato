package main

import (
	"fmt"
	"psychic-potato/linkedlist"
)

func main() {
	ReverseLinkedList()
}

func ReverseLinkedList() {
	head := linkedlist.New(7)
	fmt.Println("before reverse:")
	linkedlist.PrintList(head)

	head = linkedlist.Reverse(head)
	fmt.Println("after reverse:")
	linkedlist.PrintList(head)
}
