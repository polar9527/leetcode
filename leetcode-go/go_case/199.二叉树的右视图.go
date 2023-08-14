package go_case

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode.cn/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (65.96%)
 * Likes:    919
 * Dislikes: 0
 * Total Accepted:    307.4K
 * Total Submissions: 465.6K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1,3,4]
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,null,3]
 * 输出: [1,3]
 *
 *
 * 示例 3:
 *
 *
 * 输入: []
 * 输出: []
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [0,100]
 * -100
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
//  DFS
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		if height == len(ans) {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, height+1)
		dfs(node.Left, height+1)
	}
	dfs(root, 0)
	return ans
}

// BFS
func rightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	type i struct {
		n *TreeNode
		h int
	}
	q := []i{i{root, 0}}
	ans := []int{}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		curHeight := n.h
		if curHeight == len(ans) {
			ans = append(ans, n.n.Val)
		}
		if n.n.Right != nil {
			q = append(q, i{n.n.Right, curHeight + 1})
		}
		if n.n.Left != nil {
			q = append(q, i{n.n.Left, curHeight + 1})
		}
	}
	return ans
}

// @lc code=end
