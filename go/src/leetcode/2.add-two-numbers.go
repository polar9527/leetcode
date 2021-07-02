/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	ret := new(ListNode)
	head := ret
	total := 0
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		total += carry

		carry = total / 10

		head.Val = total % 10
		total = 0

		if l1 != nil || l2 != nil {
			head.Next = new(ListNode)
			head = head.Next
		}

	}
	if carry > 0 {
		head.Next = new(ListNode)
		head = head.Next
		head.Val = carry
	}
	return ret
}
