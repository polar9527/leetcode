package offer

import type24 "github.com/polar9527/leetcode/leetcode-go/offer/24.reverse-list/type"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *type24.ListNode) *type24.ListNode {
	var pre, post *type24.ListNode
	for head != nil {
		post = head.Next
		head.Next = pre
		pre, head = head, post
	}
	return pre
}
