package offer

import (
	type06 "github.com/polar9527/leetcode/leetcode-go/offer/06.reverse-print/type"
)

func reversePrint(head *type06.ListNode) []int {
	var count int

	return reversePrintRecursive(head, &count)
}

func reversePrintRecursive(head *type06.ListNode, count *int) []int {
	if head == nil {
		return []int{}
	}
	if *count > 10000 {
		return []int{}
	} else {
		*count += 1
		res := append(reversePrintRecursive(head.Next, count), head.Val)
		if *count > 10000 {
			return []int{}
		} else {
			return res
		}
	}
}
