package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var node *ListNode
	if l1.Val < l2.Val {
		node = l1
		node.Next = mergeTwoLists(l1.Next, l2)
	} else {
		node = l2
		node.Next = mergeTwoLists(l1, l2.Next)
	}

	return node
}
