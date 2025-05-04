package go_case

/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode.cn/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (61.07%)
 * Likes:    678
 * Dislikes: 0
 * Total Accepted:    142.3K
 * Total Submissions: 233.1K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 *
 *
 * 示例1：
 *
 *
 *
 *
 * 输入：l1 = [7,2,4,3], l2 = [5,6,4]
 * 输出：[7,8,0,7]
 *
 *
 * 示例2：
 *
 *
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[8,0,7]
 *
 *
 * 示例3：
 *
 *
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 100]
 * 0 <= node.val <= 9
 * 输入数据保证链表代表的数字无前导 0
 *
 *
 *
 *
 * 进阶：如果输入链表不能翻转该如何解决？
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	stack := func(l *ListNode) []int {
		s := []int{}
		for cur := l; cur != nil; cur = cur.Next {
			s = append(s, cur.Val)
		}
		return s
	}
	s1 := stack(l1)
	s2 := stack(l2)

	len1 := len(s1)
	len2 := len(s2)

	var res *ListNode
	carry := 0
	for len1 != 0 || len2 != 0 || carry != 0 {
		var a, b int
		if len1 != 0 {
			a = s1[len1-1]
			len1--
		} else {
			a = 0
		}
		if len2 != 0 {
			b = s2[len2-1]
			len2--
		} else {
			b = 0
		}
		sum := a + b + carry
		carry = sum / 10
		sum = sum % 10

		cur := &ListNode{Val: sum}
		cur.Next = res
		res = cur
	}
	return res
}

// @lc code=end
