/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 *
 * https://leetcode.cn/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (58.44%)
 * Likes:    1515
 * Dislikes: 0
 * Total Accepted:    615.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是 平衡二叉树
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,3,3,null,null,4,4]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在范围 [0, 5000] 内
 * -10^4 <= Node.val <= 10^4
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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var getHigh func(*TreeNode) int
	getHigh = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		leftHeight := getHigh(n.Left)
		rightHeight := getHigh(n.Right)
		if leftHeight == -1 || rightHeight == -1 {
			return -1
		}
		if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}

	return getHigh(root) != -1

}

// @lc code=end

