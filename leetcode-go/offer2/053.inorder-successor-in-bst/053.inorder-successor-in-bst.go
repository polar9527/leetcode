package oofer2

/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序搜索树
 *
 * https://leetcode.cn/problems/increasing-order-search-tree/description/
 *
 * algorithms
 * Medium (64.5%)
 * Likes:    198
 * Dislikes: 0
 * Total Accepted:    16K
 * Total Submissions: 24.7K
 *
 * 给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
 *
 * 节点 p 的后继是值比 p.val 大的节点中键值最小的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：root = [2,1,3], p = 1
 * 输出：2
 * 解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。
 * 示例 2：
 *
 *
 *
 * 输入：root = [5,3,6,2,4,null,null,1], p = 6
 * 输出：null
 * 解释：因为给出的节点没有中序后继，所以答案就返回 null 了。
 *
 *
 * 提示：
 *
 * 树中节点的数目在范围 [1, 104] 内。
 * -105 <= Node.val <= 105
 * 树中各节点的值均保证唯一。
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

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}
	node := root
	for node != nil {
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}

//  inorder
// func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
// 	stack := []*TreeNode{}
// 	var pre, cur *TreeNode = nil, root
// 	for len(stack) > 0 || cur != nil {
// 		// 用栈缓存 父节点，然后前进到二叉树最左端
// 		for cur != nil {
// 			stack = append(stack, cur)
// 			cur = cur.Left
// 		}
// 		// 中序遍历中的 visit start
// 		// 从栈中单出的这个节点，要么是最左端的叶子节点
// 		// 要么是最左端的只带右树的父节点，
// 		// 中序遍历访问父节点
// 		cur = stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		if pre == p {
// 			return cur
// 		}
// 		// 中序遍历中的 visit end
// 		//
// 		pre = cur
// 		cur = cur.Right
// 	}
// 	return nil
// }

// @lc code=end
