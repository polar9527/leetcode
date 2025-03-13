package go_case

import (
	"container/list"
	"math"
)

/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 *
 * https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (62.77%)
 * Likes:    565
 * Dislikes: 0
 * Total Accepted:    251.6K
 * Total Submissions: 400.8K
 * Testcase Example:  '[4,2,6,1,3]'
 *
 * 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
 *
 * 差值是一个正数，其数值等于两值之差的绝对值。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [4,2,6,1,3]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,0,48,null,null,12,49]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目范围是 [2, 10^4]
 * 0 <= Node.val <= 10^5
 *
 *
 *
 *
 * 注意：本题与 783
 * https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同
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
func getMinimumDifference(root *TreeNode) int {
	var preNode *TreeNode
	stack := list.New()
	sub := math.MaxInt
	stack.PushBack(root)
	for stack.Len() > 0 {
		cur := stack.Back()
		stack.Remove(cur)
		if cur.Value == nil {
			cur := stack.Back()
			stack.Remove(cur)
			node := cur.Value.(*TreeNode)
			if preNode != nil && sub > node.Val-preNode.Val {
				sub = node.Val - preNode.Val
			}
			preNode = node
		} else {
			node := cur.Value.(*TreeNode)
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
	return sub
}

// @lc code=end
