package go_case

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (77.57%)
 * Likes:    1831
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 1.7M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树 root ，返回其最大深度。
 *
 * 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,null,2]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数量在 [0, 10^4] 区间内。
 * -100 <= Node.val <= 100
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
// recursive
// func maxDepth(root *TreeNode) int {

// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}

// 	var getHeight func(root *TreeNode) int
// 	getHeight = func(root *TreeNode) int {
// 		if root == nil {
// 			return 0
// 		}
// 		return max(getHeight(root.Left), getHeight(root.Right)) + 1
// 	}

// 	return getHeight(root)
// }

func maxDepth(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	var getDepth func(*TreeNode, int)
	getDepth = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > res {
			res = depth
		}
		getDepth(root.Left, depth+1)
		getDepth(root.Right, depth+1)
		return
	}

	getDepth(root, 1)
	return res
}

// iterative
// func maxDepth(root *TreeNode) int {
// 	ans := 0
// 	if root == nil {
// 		return 0
// 	}
// 	queue := list.New()
// 	queue.PushBack(root)
// 	for queue.Len() > 0 {
// 		length := queue.Len()
// 		for i := 0; i < length; i++ {
// 			node := queue.Remove(queue.Front()).(*TreeNode)
// 			if node.Left != nil {
// 				queue.PushBack(node.Left)
// 			}
// 			if node.Right != nil {
// 				queue.PushBack(node.Right)
// 			}
// 		}
// 		ans++ //记录深度，其他的是层序遍历的板子
// 	}
// 	return ans
// }

// @lc code=end
