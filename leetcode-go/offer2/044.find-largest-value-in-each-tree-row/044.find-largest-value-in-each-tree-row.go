package offer2

import "math"

/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 *
 * https://leetcode.cn/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (66.42%)
 * Likes:    326
 * Dislikes: 0
 * Total Accepted:    126.3K
 * Total Submissions: 190.1K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
 *
 *
 *
 * 示例1：
 *
 *
 *
 *
 * 输入: root = [1,3,2,5,3,null,9]
 * 输出: [1,3,9]
 *
 *
 * 示例2：
 *
 *
 * 输入: root = [1,2,3]
 * 输出: [1,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 二叉树的节点个数的范围是 [0,10^4]
 * -2^31 <= Node.val <= 2^31 - 1
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func largestValues(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		maxVal := math.MinInt32
		level := q
		q = nil
		for _, n := range level {
			maxVal = max(maxVal, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		ans = append(ans, maxVal)
	}
	return ans
}

// DFS
func largestValuesDFS(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curHeight int) {
		if node == nil {
			return
		}
		if curHeight == len(ans) {
			ans = append(ans, node.Val)
		} else {
			ans[curHeight] = max(ans[curHeight], node.Val)
		}
		dfs(node.Left, curHeight+1)
		dfs(node.Right, curHeight+1)
	}
	dfs(root, 0)
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// func max(a, b int) int {
// 	if b > a {
// 		return b
// 	}
// 	return a
// }

// @lc code=end
