package go_case

/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 *
 * https://leetcode.com/problems/house-robber-iii/description/
 *
 * algorithms
 * Medium (54.24%)
 * Likes:    8436
 * Dislikes: 140
 * Total Accepted:    378.8K
 * Total Submissions: 698.4K
 * Testcase Example:  '[3,2,3,null,3,null,1]'
 *
 * The thief has found himself a new place for his thievery again. There is
 * only one entrance to this area, called root.
 *
 * Besides the root, each house has one and only one parent house. After a
 * tour, the smart thief realized that all houses in this place form a binary
 * tree. It will automatically contact the police if two directly-linked houses
 * were broken into on the same night.
 *
 * Given the root of the binary tree, return the maximum amount of money the
 * thief can rob without alerting the police.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,2,3,null,3,null,1]
 * Output: 7
 * Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [3,4,5,1,3,null,1]
 * Output: 9
 * Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * 0 <= Node.val <= 10^4
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var robTree func(*TreeNode) []int
	robTree = func(cur *TreeNode) []int {
		if cur == nil {
			return []int{0, 0}
		}
		// 后序遍历
		left := robTree(cur.Left)
		right := robTree(cur.Right)

		// 考虑去偷当前的屋子
		robCur := cur.Val + left[0] + right[0]
		// 考虑不去偷当前的屋子
		notRobCur := max(left[0], left[1]) + max(right[0], right[1])

		// 注意顺序：0:不偷，1:去偷
		return []int{notRobCur, robCur}
	}

	res := robTree(root)
	return max(res[0], res[1])
}

// @lc code=end
