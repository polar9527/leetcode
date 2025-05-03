package go_case

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode.cn/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (72.29%)
 * Likes:    2214
 * Dislikes: 0
 * Total Accepted:    864.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 100] 内
 * 0 <= Node.val <= 100
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
// func swapPairs(head *ListNode) *ListNode {

// 	dummy := &ListNode{Next: head}
// 	cur := dummy
// 	for cur.Next != nil && cur.Next.Next != nil {
// 		// 指向第 1 个节点
// 		tmp1 := cur.Next
// 		// 指向第 3 个节点
// 		tmp2 := cur.Next.Next.Next

// 		// 交换
// 		cur.Next = cur.Next.Next
// 		cur.Next.Next = tmp1
// 		cur.Next.Next.Next = tmp2

// 		// 移动指针
// 		cur = cur.Next.Next
// 	}

// 	return dummy.Next

// }

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		// 指向第 1 个节点
		node1 := cur.Next
		// 指向第 2 个节点
		node2 := cur.Next.Next
		// 指向第 3 个节点
		node3 := cur.Next.Next.Next

		// 交换
		cur.Next = node2
		cur.Next.Next = node1
		cur.Next.Next.Next = node3

		// 移动指针
		cur = cur.Next.Next
	}
	return dummy.Next
}

// func swapPairs(head *ListNode) *ListNode {
// 	n := 0
// 	for cur := head; cur != nil; cur = cur.Next {
// 		n++
// 	}

// 	dummy := &ListNode{Next: head}

// 	start := dummy
// 	var cur, pre *ListNode = head, nil
// 	for ; n >= 2; n -= 2 {
// 		for i := 0; i < 2; i++ {
// 			next := cur.Next
// 			cur.Next = pre
// 			pre = cur
// 			cur = next
// 		}
// 		nextStart := start.Next
// 		start.Next.Next = cur
// 		start.Next = pre
// 		start = nextStart
// 	}
// 	return dummy.Next
// }

// @lc code=end
