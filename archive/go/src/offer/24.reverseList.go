package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var node, preNode, nextNode *ListNode
	node = head
	nextNode = node.Next

	for node != nil {
		// reverse
		node.Next = preNode

		// step forward
		// step 1
		preNode = node
		// step 2
		node = nextNode
		// step 3
		if nextNode != nil {
			nextNode = node.Next
		}
	}
	return preNode
}
