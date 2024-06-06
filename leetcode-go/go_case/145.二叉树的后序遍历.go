package go_case

import "container/list"

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (76.59%)
 * Likes:    1176
 * Dislikes: 0
 * Total Accepted:    781.4K
 * Total Submissions: 1M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[3,2,1]
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
 * 树中节点的数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
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
//  recursive
// func postorderTraversal(root *TreeNode) []int {
// 	var res []int
// 	if root == nil {
// 		return res
// 	}

// 	var traveler func(*TreeNode)
// 	traveler = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}

// 		traveler(node.Left)
// 		traveler(node.Right)
// 		res = append(res, node.Val)
// 	}
// 	traveler(root)
// 	return res
// }

// iterative
// func postorderTraversal(root *TreeNode) []int {
// 	var res []int
// 	if root == nil {
// 		return res
// 	}
// 	reverse := func(arr []int) {
// 		l, r := 0, len(arr)-1
// 		for l < r {
// 			arr[l], arr[r] = arr[r], arr[l]
// 			l++
// 			r--
// 		}
// 	}

// 	var stack []*TreeNode
// 	stack = append(stack, root)
// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		if node == nil {
// 			continue
// 		}
// 		res = append(res, node.Val)
// 		stack = append(stack, node.Left)
// 		stack = append(stack, node.Right)
// 	}
// 	reverse(res)
// 	return res

// }

func postorderTraversal(root *TreeNode) []int {
	// uniform iterative
	// 后序  左右中
	// 压栈  中右左
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
			stack.PushBack(node)
			stack.PushBack(nil)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		}
	}
	return res
}

// @lc code=end
