/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode.cn/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (57.18%)
 * Likes:    1961
 * Dislikes: 0
 * Total Accepted:    629.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目为 n
 * 1
 * -500
 * 1
 *
 *
 *
 *
 * 进阶： 你可以使用一趟扫描完成反转吗？
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}

	// 移动到 left 之前的一个节点，
	// 当 left == 1 时，start == head
	start := dummy
	for i := 1; i <= left-1; i++ {
		start = start.Next
	}

	var cur, pre *ListNode = start.Next, nil

	for i := 1; i <= right-left+1; i++ {
		next := cur.Next
		// 反转节点
		cur.Next = pre
		// 前进一位
		pre = cur
		cur = next
	}
	// 将反转链表重新接入
	start.Next.Next = cur
	start.Next = pre

	return dummy.Next
}

// @lc code=end

