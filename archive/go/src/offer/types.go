package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
