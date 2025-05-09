package go_case

import "container/list"

/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 *
 * https://leetcode.cn/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Easy (81.66%)
 * Likes:    1136
 * Dislikes: 0
 * Total Accepted:    397.7K
 * Total Submissions: 487K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
 *
 * 完全二叉树
 * 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h
 * 层，则该层包含 1~ 2^h 个节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,4,5,6]
 * 输出：6
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目范围是[0, 5 * 10^4]
 * 0
 * 题目数据保证输入的树是 完全二叉树
 *
 *
 *
 *
 * 进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？
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

// recursive
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	var right, left int
// 	if root.Right != nil {
// 		right = countNodes(root.Right)
// 	}
// 	if root.Left != nil {
// 		left = countNodes(root.Left)
// 	}
// 	return right + left + 1
// }

// iterative
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New()
	q.PushBack(root)
	res := 0
	for q.Len() > 0 {
		n := q.Len()
		for i := 0; i < n; i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
			res++
		}
	}
	return res
}

// @lc code=end
