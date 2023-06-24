/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// b >= a >= 0, a and b belong to int
// i is the point where tortoise and hare meet
// at the moment they meet, hare stepped 2i
// u is the lenght of tail in list which has a cycle
// lamada is the lenght of the cycle
// i  = u + a*lamada
// 2i = u + b*lamada
// X[i + u] = X[(b - a)*lamada + u] = X[u]
func detectCycle(head *ListNode) *ListNode {
	var cycPoint *ListNode
	if head == nil {
		return cycPoint
	}
	t, h := head, head
	started := false
	hasCycle := false
	for h != nil && h.Next != nil {
		if t == h {
			if started {
				hasCycle = true
				break
			} else {
				started = true
			}
		}
		t = t.Next
		h = h.Next.Next
	}

	if hasCycle {
		h = head
		for h != nil && t != nil {
			if h == t {
				cycPoint = h
				break
			}
			h = h.Next
			t = t.Next
		}
	}

	return cycPoint
}
