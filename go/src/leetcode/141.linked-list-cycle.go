/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
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
// Both t and h start at head,
// the reason of doing this can prepare the start point of the second round in problem 142
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	t, h := head, head
	started := false
	for h != nil && h.Next != nil {
		if t == h {
			if started {
				return true
			} else {
				started = true
			}
		}
		t = t.Next
		h = h.Next.Next

	}

	return false
}
