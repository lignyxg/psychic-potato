package linkedlist

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	val  int
	next *ListNode
}

// 返回长度为n的一个链表
func New(num int) *ListNode {
	root := &ListNode{
		val:  rand.Intn(100),
		next: nil,
	}

	for i := 0; i < num; i++ {
		node := new(ListNode)
		node.val = rand.Intn(100) + 1
		node.next = root.next
		root.next = node
	}
	return root
}

func PrintList(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf("%d->", p.val)
		p = p.next
	}
	fmt.Println("nil")
}

//反转单链表
func Reverse(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	p := head.next
	newHead := Reverse(p)
	p.next = head
	head.next = nil // 如果没有这一步,会在新链表最后两个元素生成一个环
	return newHead
}
