/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return head
	}

	fast, slow := head, head

	for i := 0; i < n; i++ {
		if fast.Next == nil {
			return head.Next
		}
		fast = fast.Next
	}
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next

	return head
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	lenght := 0
// 	pList := head
// 	for pList != nil {
// 		lenght++
// 		pList = pList.Next
// 	}
// 	dIndex := lenght - n
// 	if dIndex == 0 {
// 		head = head.Next
// 	} else {
// 		pList = head
// 		for i := 1; i < dIndex; i++ {
// 			pList = pList.Next
// 		}
// 		pList.Next = pList.Next.Next

// 	}
// 	return head
// }
