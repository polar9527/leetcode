package go_case

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 *
 * https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (45.76%)
 * Likes:    2614
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.6M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1], n = 1
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中结点的数目为 sz
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 *
 * 进阶：你能尝试使用一趟扫描实现吗？
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, dummy
	for i := n; i > 0 && fast.Next != nil; i-- {
		// 往前走 n 个节点
		fast = fast.Next
	}
	// 此时 slow 和 fast 之间,包括 slow 和 fast 指向的节点，一共有 n+1 个
	for fast.Next != nil {
		// fast 走到最后一个节点停止
		// 此时slow指向倒数第 n+1 个节点，
		// 或者 如果 list 的长度 == n, 即 fast 和 slow 之间没有 n+1个节点, 只有 n 个节点
		// 此循环直接跳出，slow 仍然指向 dummy，
		fast = fast.Next
		slow = slow.Next
	}
	// 此时 slow 指向倒数第 n-1 个 node
	// remove the Nth node

	// 题目中的条件是 至少有一个节点，所以不用担心 slow.Next 为 nil
	// 当 slow 仍然指向 dummy 的时候，意味着 list 的长度刚好为 n, 因此删除的是 slow.Next == head 节点
	slow.Next = slow.Next.Next
	return dummy.Next

}

// @lc code=end
