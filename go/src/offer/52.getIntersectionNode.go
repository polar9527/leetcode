package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == headB {
		return headA
	}

	al := lengthOfList(headA)
	bl := lengthOfList(headB)

	var longList, shortList *ListNode
	step := al - bl
	if step < 0 {
		step = -step
		longList = headB
		shortList = headA
	} else {
		longList = headA
		shortList = headB
	}

	for i := 0; i < step; i++ {
		longList = longList.Next
	}
	for longList != nil && shortList != nil {
		if longList == shortList {
			return longList
		}
		if longList.Next == shortList.Next {
			return longList.Next
		}
		longList = longList.Next
		shortList = shortList.Next
	}
	return nil
}

func lengthOfList(node *ListNode) int {
	count := 0
	if node == nil {
		return count
	}
	for node != nil {
		count++
		node = node.Next
	}
	return count
}
