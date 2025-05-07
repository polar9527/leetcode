package go_case

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode.cn/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (53.35%)
 * Likes:    1746
 * Dislikes: 0
 * Total Accepted:    614.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 206
// func reverseList(head *ListNode) *ListNode {
// 	var pre, post *ListNode
// 	for head != nil {
// 		post, head.Next = head.Next, pre
// 		pre, head = head, post
// 	}
// 	return pre
// }

// 876
// func middleNode(head *ListNode) *ListNode {
// 	slow, fast := head, head
// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	return slow
// }

// func isPalindromeList(head *ListNode) bool {
// 	if head == nil {
// 		return true
// 	}

// 	// 找到前半部分链表的尾节点并反转后半部分链表
// 	firstHalfEnd := middleNode(head)
// 	secondHalfStart := reverseList(firstHalfEnd)

// 	// 判断是否回文
// 	p1 := head
// 	p2 := secondHalfStart
// 	result := true
// 	for result && p2 != nil {
// 		if p1.Val != p2.Val {
// 			result = false
// 		}
// 		p1 = p1.Next
// 		p2 = p2.Next
// 	}

// 	// 还原链表并返回结果
// 	firstHalfEnd.Next = reverseList(secondHalfStart)
// 	return result
// }

func isPalindrome(head *ListNode) bool {

	middleNode := func(head *ListNode) *ListNode {
		// 偶数
		// 1-2-3-4-nil
		// slow -> 3; fast -> nil

		// 奇数
		// 1-2-3-4-5
		// slow -> 3; fast -> 5

		slow := head
		fast := head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}

		return slow
	}

	reverseList := func(head *ListNode) *ListNode {
		var pre *ListNode
		cur := head
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		return pre
	}

	mid := middleNode(head)
	secondHalfStart := reverseList(mid)

	l1 := head
	l2 := secondHalfStart
	result := true
	for result && l2 != nil {
		// 此时  nil
		//       ^
		//       |
		// 1->2->3<-4<-5
		// 或者  nil
		//       ^
		//       |
		// 1->2->3<-4
		// 所以只需要判断 l2 != nil
		if l1.Val != l2.Val {
			result = false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	reverseList(secondHalfStart)
	return result
}

// @lc code=end
