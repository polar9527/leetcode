package go_case

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode.cn/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (63.26%)
 * Likes:    1110
 * Dislikes: 0
 * Total Accepted:    411.5K
 * Total Submissions: 650.4K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：[[5,4,11,2],[5,8,4,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,3], targetSum = 5
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], targetSum = 0
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点总数在范围 [0, 5000] 内
 * -1000
 * -1000
 *
 *
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, targetSum int, path []int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if targetSum == node.Val {
				path = append(path, node.Val)
				res = append(res, append([]int{}, path...))
				return
			}
		}

		dfs(node.Left, targetSum-node.Val, append(path, node.Val))
		dfs(node.Right, targetSum-node.Val, append(path, node.Val))

	}

	dfs(root, targetSum, []int{})
	return res
}

// @lc code=end
