package go_case

import "container/list"

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Easy (76.78%)
 * Likes:    2084
 * Dislikes: 0
 * Total Accepted:    1.5M
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
// func inorderTraversal(root *TreeNode) []int {
// 	var res []int
// 	if root == nil {
// 		return []int{}
// 	}
// 	var traveler func(*TreeNode)
// 	traveler = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		traveler(node.Left)
// 		res = append(res, node.Val)
// 		traveler(node.Right)
// 	}
// 	traveler(root)
// 	return res
// }

// iterative
// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	if root == nil {
// 		return res
// 	}

// 	stack := []*TreeNode{}
// 	cur := root
// 	for cur != nil || len(stack) > 0 {
// 		if cur != nil {
// 			stack = append(stack, cur)
// 			cur = cur.Left
// 		} else {
// 			cur = stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			res = append(res, cur.Val)
// 			cur = cur.Right

// 		}
// 	}
// 	return res
// }

func inorderTraversal(root *TreeNode) []int {
	// uniform iterative
	// 中序  左中右
	// 压栈  右中左
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	var res []int
	var node *TreeNode
	for stack.Len() > 0 {
		cur := stack.Back()
		stack.Remove(cur)
		if cur.Value == nil {
			cur = stack.Back()
			stack.Remove(cur)
			node = cur.Value.(*TreeNode)
			res = append(res, node.Val)
			continue
		} else {
			node = cur.Value.(*TreeNode)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			stack.PushBack(node)
			stack.PushBack(nil)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		}
	}
	return res
}

// @lc code=end
