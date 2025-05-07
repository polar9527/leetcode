package go_case

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 *
 * https://leetcode.cn/problems/reorder-list/description/
 *
 * algorithms
 * Medium (64.95%)
 * Likes:    1337
 * Dislikes: 0
 * Total Accepted:    277.6K
 * Total Submissions: 422.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * 请将其重新排列后变为：
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[1,4,2,3]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[1,5,2,4,3]
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 5 * 10^4]
 * 1 <= node.val <= 1000
 *
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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l2 := reverseList(mid)
	// mergeList(head, l2)
	l1 := head
	for l2.Next != nil {
		nextL1 := l1.Next
		nextL2 := l2.Next

		l1.Next = l2
		l2.Next = nextL1

		l1 = nextL1
		l2 = nextL2
	}

}

// 876
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 206
func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	cur := head
	for cur != nil {
		// 预存下一个节点指针 next
		// 反转 节点 head.Next，指向上一个节点 pre
		next, cur.Next = cur.Next, pre
		// 上一个节点 pre 前进
		// 当前节点 head 前进
		pre, cur = cur, next
	}
	return pre
}

// @lc code=end
