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

	node := head
	var nextNode, preNode *ListNode

	for node != nil {
		// reverse
		nextNode = node.Next
		node.Next = preNode
		// step
		preNode = node
		node = nextNode
		if nextNode != nil {
			nextNode = nextNode.Next
		}
	}

	return preNode
}
