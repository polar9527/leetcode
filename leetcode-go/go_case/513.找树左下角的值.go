package go_case

/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
 *
 * https://leetcode.cn/problems/find-bottom-left-tree-value/description/
 *
 * algorithms
 * Medium (73.41%)
 * Likes:    508
 * Dislikes: 0
 * Total Accepted:    200.3K
 * Total Submissions: 272.9K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
 *
 * 假设二叉树中至少有一个节点。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: root = [2,1,3]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * ⁠
 *
 *
 * 输入: [1,2,3,4,null,5,6,null,null,7]
 * 输出: 7
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [1,10^4]
 * -2^31
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

// DFS
// func findBottomLeftValue(root *TreeNode) (curVal int) {
// 	if root == nil {
// 		return 0
// 	}
// 	var res int
// 	maxDepth := math.MinInt
// 	var dfs func(root *TreeNode, curDepth int)
// 	dfs = func(root *TreeNode, curDepth int) {
// 		if root == nil {
// 			return
// 		}
// 		if root.Left == nil && root.Right == nil {

// 			if curDepth > maxDepth {
// 				maxDepth = curDepth
// 				res = root.Val
// 			}
// 			return
// 		}
// 		dfs(root.Left, curDepth+1)
// 		dfs(root.Right, curDepth+1)
// 	}

// 	dfs(root, 0)
// 	return res
// }

// BFS
// func findBottomLeftValue(root *TreeNode) int {
// 	var gradation int
// 	queue := list.New()

// 	queue.PushBack(root)
// 	for queue.Len() > 0 {
// 		length := queue.Len()
// 		for i := 0; i < length; i++ {
// 			node := queue.Remove(queue.Front()).(*TreeNode)
// 			if i == 0 {
// 				gradation = node.Val
// 			}
// 			if node.Left != nil {
// 				queue.PushBack(node.Left)
// 			}
// 			if node.Right != nil {
// 				queue.PushBack(node.Right)
// 			}
// 		}
// 	}
// 	return gradation
// }

func findBottomLeftValue(root *TreeNode) int {
	res := 0

	q := []*TreeNode{root}

	for len(q) > 0 {
		amoutl := len(q)

		for i := 0; i < amoutl; i++ {
			node := q[0]
			q = q[1:]
			if i == 0 {
				res = node.Val
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return res
}

// @lc code=end
