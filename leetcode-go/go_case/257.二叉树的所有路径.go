package go_case

import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode.cn/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (70.74%)
 * Likes:    1140
 * Dislikes: 0
 * Total Accepted:    398.1K
 * Total Submissions: 562.7K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,null,5]
 * 输出：["1->2->5","1->3"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：["1"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 100] 内
 * -100 <= Node.val <= 100
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
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	var backtracing func(*TreeNode, string)
	backtracing = func(node *TreeNode, path string) {

		if node.Left == nil && node.Right == nil {
			v := path + strconv.Itoa(node.Val)
			res = append(res, v)
			return
		}

		path = path + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			backtracing(node.Left, path)
		}

		if node.Right != nil {
			backtracing(node.Right, path)
		}
	}
	backtracing(root, "")
	return res
}

// @lc code=end
