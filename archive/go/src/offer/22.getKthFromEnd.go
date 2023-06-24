package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}

	first := head
	second := first

	for first != nil {
		first = first.Next
		if k > 0 {
			k--
		} else {
			second = second.Next
		}
	}

	if k != 0 {
		return nil
	}
	return second
}
