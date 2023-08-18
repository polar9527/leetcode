package go_case

/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入二叉搜索树
 *
 * https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (63.51%)
 * Likes:    482
 * Dislikes: 0
 * Total Accepted:    114.8K
 * Total Submissions: 180.8K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * 给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: root = [5,3,6,2,4,null,7], k = 9
 * 输出: true
 *
 *
 * 示例 2：
 *
 *
 * 输入: root = [5,3,6,2,4,null,7], k = 28
 * 输出: false
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是  [1, 10^4].
 * -10^4 <= Node.val <= 10^4
 * 题目数据保证，输入的 root 是一棵 有效 的二叉搜索树
 * -10^5 <= k <= 10^5
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

// Two Point
func findTarget(root *TreeNode, k int) bool {
	left, right := root, root
	leftStk := []*TreeNode{left}
	for left.Left != nil {
		leftStk = append(leftStk, left.Left)
		left = left.Left
	}
	rightStk := []*TreeNode{right}
	for right.Right != nil {
		rightStk = append(rightStk, right.Right)
		right = right.Right
	}
	for left != right {
		sum := left.Val + right.Val
		if sum == k {
			return true
		}
		if sum < k {
			left = leftStk[len(leftStk)-1]
			leftStk = leftStk[:len(leftStk)-1]
			for node := left.Right; node != nil; node = node.Left {
				leftStk = append(leftStk, node)
			}
		} else {
			right = rightStk[len(rightStk)-1]
			rightStk = rightStk[:len(rightStk)-1]
			for node := right.Left; node != nil; node = node.Right {
				rightStk = append(rightStk, node)
			}
		}
	}
	return false
}

// DFS
// func findTarget(root *TreeNode, k int) bool {
// 	set := map[int]struct{}{}
// 	var dfs func(*TreeNode)
// 	dfs = func(n *TreeNode) {
// 		if node == nil {
// 			return false
// 		}
// 		if _, ok := set[k-set[n.Val]]; ok {
// 			return true
// 		}
// 		set[n.Val] = struct{}{}
// 		return dfs(n.Left) || dfs(n.Right)
// 	}
// 	return dfs(root)
// }

// BFS
// func findTarget(root *TreeNode, k int) bool {
// 	set := map[int]struct{}{}
// 	q := []*TreeNode{root}
// 	for len(q) > 0 {
// 		n := q[0]
// 		q := q[1:]
// 		if _, ok := set[k-n.Val]; ok {
// 			return true
// 		}
// 		set[node.Val] = struct{}{}
// 		if node.Left != nil {
// 			q = append(q, node.Left)
// 		}
// 		if node.Right != nil {
// 			q = append(q, node.Right)
// 		}
// 	}
// }

// @lc code=end
