/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func isPalindrome(head *ListNode) bool {
// 	if head == nil || head.Next == nil {
// 		return true
// 	}

// 	s, d := head, head
// 	var p1, p2 *ListNode

// 	for {
// 		d = d.Next.Next

// 		p1 = s
// 		s = s.Next
// 		p1.Next = p2
// 		p2 = p1

// 		if d == nil || d.Next == nil {
// 			break
// 		}
// 	}

// 	if d != nil {
// 		s = s.Next
// 	}

// 	left, right := p1, s

// 	for left != nil && right != nil {
// 		if left.Val != right.Val {
// 			return false
// 		}
// 		left, right = left.Next, right.Next
// 	}
// 	return true
// }

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	s := head
	d := head.Next
	for d != nil && d.Next != nil {
		d = d.Next.Next
		s = s.Next
	}

	tail := reverseList(s)

	for tail != nil && head != nil {
		if tail.Val != head.Val {
			return false
		}
		tail, head = tail.Next, head.Next
	}
	return true

}

func reverseList(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		head.Next, tail, head = tail, head, head.Next
	}
	return tail
}
