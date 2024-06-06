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

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if root == nil {
		return true
	}

	var getDepth func(*TreeNode, int) int
	getDepth = func(root *TreeNode, depth int) int {
		if root == nil {
			return 0
		}

		ld := getDepth(root.Left, depth+1)
		rd := getDepth(root.Right, depth+1)
		if ld == -1 || rd == -1 {
			return -1
		}
		diff := ld - rd
		if diff > 1 || diff < -1 {
			return -1
		}

		return max(ld, rd) + 1
	}
	return getDepth(root, 1) != -1
}

// @lc code=end

