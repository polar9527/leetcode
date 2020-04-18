/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (47.60%)
 * Likes:    1503
 * Dislikes: 40
 * Total Accepted:    384.3K
 * Total Submissions: 788.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its level order traversal as:
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
	dqueueLevel := make([][]*TreeNode, 0)
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	currentNode := root
	dqueueLevel = append(dqueueLevel, []*TreeNode{currentNode})
	currentLevel := dqueueLevel[0]
	for len(dqueueLevel) > 0 {
		currentLevel = dqueueLevel[0]
		dqueueLevel = dqueueLevel[1:]
		var tempLevel []*TreeNode
		var tempLevelVal []int
		for _, n := range currentLevel {
			tempLevelVal = append(tempLevelVal, n.Val)
			if n.Left != nil {
				tempLevel = append(tempLevel, n.Left)
			}
			if n.Right != nil {
				tempLevel = append(tempLevel, n.Right)
			}
		}
		if len(tempLevel) > 0 {
			dqueueLevel = append(dqueueLevel, tempLevel)
		}
		if len(tempLevelVal) > 0 {
			ret = append(ret, tempLevelVal)
		}
	}
	return ret
}

// @lc code=end
