package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func (l *List) addNode(val int) {
	newNode := &ListNode{val, nil}
	if l.head == nil {
		l.head = newNode
	}
	if l.tail == nil {
		l.tail = newNode
		return
	}
	l.tail.Next = newNode
	l.tail = newNode
}

func main() {
	l := new(List)
	l.addNode(1)
	l.addNode(2)
	l.addNode(3)
	l.addNode(4)
	l.addNode(5)
	l.addNode(6)

	s := reversePrint(l.head)
	fmt.Println(s)
}

func reversePrint(head *ListNode) []int {
	cur := head
	var stack []int
	if head == nil {
		return stack
	}
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	// reverse array
	for lo, hi := 0, len(stack)-1; lo < hi; {
		stack[lo], stack[hi] = stack[hi], stack[lo]
		lo++
		hi--
	}
	return stack
}
