/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (69.23%)
 * Likes:    1028
 * Dislikes: 0
 * Total Accepted:    261.9K
 * Total Submissions: 377.1K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 *
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var tail *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = tail

		tail = cur
		cur = next
	}
	return tail
}

func reverseList1(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		head.Next, tail, head = tail, head, head.Next
	}
	return tail
}

// @lc code=end
