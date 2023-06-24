package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	preNode := head
	curNode := head

	for curNode != nil {
		if curNode.Val == val {
			if curNode == head {
				head = head.Next
			} else {
				preNode.Next = curNode.Next
			}
			break
		} else {
			preNode = curNode
			curNode = curNode.Next
		}
	}

	return head
}
