package go_case

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (67.25%)
 * Likes:    1961
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.6M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[9,20],[15,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 2000] 内
 * -1000 <= Node.val <= 1000
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
// func levelOrder(root *TreeNode) [][]int {

// 	var res [][]int

// 	var traveler func(*TreeNode, int)
// 	traveler = func(root *TreeNode, depth int) {
// 		if root == nil {
// 			return
// 		}

// 		// at the first time we arrive the Nth level,
// 		// we need to create a slice for this level to load the values
// 		if len(res) == depth {
// 			res = append(res, []int{})
// 		}

// 		res[depth] = append(res[depth], root.Val)

// 		traveler(root.Left, depth+1)
// 		traveler(root.Right, depth+1)
// 	}
// 	traveler(root, 0)
// 	return res
// }

// iterative
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		amountOfEachLevel := len(queue)
		level := []int{}
		for i := 0; i < amountOfEachLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// @lc code=end
