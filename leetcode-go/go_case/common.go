package go_case

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val  int
	Next *ListNode
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
