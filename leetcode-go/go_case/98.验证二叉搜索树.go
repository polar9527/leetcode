package go_case

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode.cn/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (37.96%)
 * Likes:    2344
 * Dislikes: 0
 * Total Accepted:    920.7K
 * Total Submissions: 2.4M
 * Testcase Example:  '[2,1,3]'
 *
 * 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
 *
 * 有效 二叉搜索树定义如下：
 *
 *
 * 节点的左子树只包含 小于 当前节点的数。
 * 节点的右子树只包含 大于 当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [2,1,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,1,4,null,null,3,6]
 * 输出：false
 * 解释：根节点的值是 5 ，但是右子节点的值是 4 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围在[1, 10^4] 内
 * -2^31 <= Node.val <= 2^31 - 1
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
// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return false
// 	}
// 	queue := list.New()
// 	var dfs func(*TreeNode)
// 	dfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		if node.Left == nil && node.Right == nil {
// 			queue.PushBack(node)
// 			return
// 		}
// 		dfs(node.Left)
// 		queue.PushBack(node)
// 		dfs(node.Right)
// 	}

// 	dfs(root)

// 	preVal := queue.Front().Value.(*TreeNode).Val
// 	queue.Remove(queue.Front())
// 	for queue.Len() > 0 {
// 		if preVal >= queue.Front().Value.(*TreeNode).Val {
// 			return false
// 		}

// 		preVal = queue.Front().Value.(*TreeNode).Val
// 		queue.Remove(queue.Front())
// 	}

//		return true
//	}
//
// recursive
// func isValidBST(root *TreeNode) bool {
// 	if root == nil {
// 		return false
// 	}

// 	var dfs func(*TreeNode) bool
// 	max := math.MinInt
// 	dfs = func(n *TreeNode) bool {
// 		if n == nil {
// 			return true
// 		}

// 		if !dfs(n.Left) {
// 			return false
// 		}

// 		if n.Val > max {
// 			max = n.Val
// 		} else {
// 			return false
// 		}
// 		if !dfs(n.Right) {
// 			return false
// 		}
// 		return true
// 	}
// 	return dfs(root)
// }

// iterative
func isValidBST(root *TreeNode) bool {

	if root == nil {
		return false
	}
	stack := list.New()
	stack.PushBack(root)
	var node, preNode *TreeNode
	for stack.Len() > 0 {
		cur := stack.Back()
		stack.Remove(cur)
		if cur.Value == nil {
			cur = stack.Back()
			stack.Remove(cur)
			node = cur.Value.(*TreeNode)
			if preNode != nil && preNode.Val >= node.Val {
				return false
			}
			preNode = node
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

	return true
}

// @lc code=end
