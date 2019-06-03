/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		head.Next, tail, head = tail, head, head.Next
	}
	return tail
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	ret := head
// 	node := &ListNode{}
// 	for head.Next != nil {
// 		if head == ret {
// 			head = head.Next
// 			ret.Next = nil
// 			continue
// 		}
// 		node = head
// 		head = head.Next

// 		node.Next = ret
// 		ret = node
// 	}
// 	if head == ret {
// 		return head
// 	}
// 	head.Next = ret
// 	ret = head
// 	return ret
// }
