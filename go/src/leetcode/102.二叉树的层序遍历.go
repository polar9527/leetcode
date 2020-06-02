/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (61.71%)
 * Likes:    516
 * Dislikes: 0
 * Total Accepted:    140.4K
 * Total Submissions: 223.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层次遍历结果：
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 *
 *
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	levels := make([][]*TreeNode, 0)
	levels = append(levels, []*TreeNode{root})
	for len(levels) != 0 {
		curLevel := levels[0]
		levels = levels[1:]
		newLevel := make([]*TreeNode, 0)
		an := make([]int, 0)
		for _, node := range curLevel {
			an = append(an, node.Val)
			if node.Left != nil {
				newLevel = append(newLevel, node.Left)
			}
			if node.Right != nil {
				newLevel = append(newLevel, node.Right)
			}
		}
		ans = append(ans, an)
		if len(newLevel) != 0 {
			levels = append(levels, newLevel)
		}
	}
	return ans
}

// @lc code=end
