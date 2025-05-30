/*
 * @lc app=leetcode.cn id=3217 lang=golang
 *
 * [3217] 从链表中移除在数组中存在的节点
 *
 * https://leetcode.cn/problems/delete-nodes-from-linked-list-present-in-array/description/
 *
 * algorithms
 * Medium (62.47%)
 * Likes:    13
 * Dislikes: 0
 * Total Accepted:    12.3K
 * Total Submissions: 19.4K
 * Testcase Example:  '[1,2,3]\n[1,2,3,4,5]'
 *
 * 给你一个整数数组 nums 和一个链表的头节点 head。从链表中移除所有存在于 nums 中的节点后，返回修改后的链表的头节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： nums = [1,2,3], head = [1,2,3,4,5]
 *
 * 输出： [4,5]
 *
 * 解释：
 *
 *
 *
 * 移除数值为 1, 2 和 3 的节点。
 *
 *
 * 示例 2：
 *
 *
 * 输入： nums = [1], head = [1,2,1,2,1,2]
 *
 * 输出： [2,2,2]
 *
 * 解释：
 *
 *
 *
 * 移除数值为 1 的节点。
 *
 *
 * 示例 3：
 *
 *
 * 输入： nums = [5], head = [1,2,3,4]
 *
 * 输出： [1,2,3,4]
 *
 * 解释：
 *
 *
 *
 * 链表中不存在值为 5 的节点。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^5
 * nums 中的所有元素都是唯一的。
 * 链表中的节点数在 [1, 10^5] 的范围内。
 * 1 <= Node.val <= 10^5
 * 输入保证链表中至少有一个值没有在 nums 中出现过。
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
func modifiedList(nums []int, head *ListNode) *ListNode {
	mn := make(map[int]bool)
	for i := range nums {
		mn[nums[i]] = true
	}

	dummy := &ListNode{Next: head}

	pre := dummy
	cur := head

	for cur != nil {
		if mn[cur.Val] {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}

// @lc code=end

