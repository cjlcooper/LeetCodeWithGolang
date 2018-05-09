package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l1.Val = 2
	l1.Next = l2
	fmt.Println(l1)
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

// }