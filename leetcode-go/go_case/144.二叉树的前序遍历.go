package go_case

import "container/list"

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (71.87%)
 * Likes:    1251
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,2,3]
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
 * 示例 4：
 *
 *
 * 输入：root = [1,2]
 * 输出：[1,2]
 *
 *
 * 示例 5：
 *
 *
 * 输入：root = [1,null,2]
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100
 *
 *
 *
 *
 * 进阶：递归算法很简单，你可以通过迭代算法完成吗？
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
// func preorderTraversal(root *TreeNode) []int {

// 	var res []int

// 	var travler func(*TreeNode)
// 	travler = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		res = append(res, node.Val)
// 		travler(node.Left)
// 		travler(node.Right)
// 	}
// 	travler(root)
// 	return res
// }

// iterative
// func preorderTraversal(root *TreeNode) []int {

// 	var res []int
// 	var stack []*TreeNode

// 	stack = append(stack, root)
// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		if node == nil {
// 			continue
// 		}
// 		res = append(res, node.Val)
// 		stack = append(stack, node.Right)
// 		stack = append(stack, node.Left)
// 	}
// 	return res
// }

func preorderTraversal(root *TreeNode) []int {
	// uniform iterative
	// 前序  中左右
	// 压栈  右左中
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
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			stack.PushBack(node)
			stack.PushBack(nil)
		}
	}
	return res
}

// @lc code=end
