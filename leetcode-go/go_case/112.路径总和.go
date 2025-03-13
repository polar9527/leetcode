package go_case

/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 *
 * https://leetcode.cn/problems/path-sum/description/
 *
 * algorithms
 * Easy (54.30%)
 * Likes:    1357
 * Dislikes: 0
 * Total Accepted:    685.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点
 * 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
 * 输出：true
 * 解释：等于目标和的根节点到叶节点路径如上图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,3], targetSum = 5
 * 输出：false
 * 解释：树中存在两条根节点到叶子节点的路径：
 * (1 --> 2): 和为 3
 * (1 --> 3): 和为 4
 * 不存在 sum = 5 的根节点到叶子节点的路径。
 *
 * 示例 3：
 *
 *
 * 输入：root = [], targetSum = 0
 * 输出：false
 * 解释：由于树是空的，所以不存在根节点到叶子节点的路径。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [0, 5000] 内
 * -1000 <= Node.val <= 1000
 * -1000 <= targetSum <= 1000
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
//  DFS
// func hasPathSum(root *TreeNode, targetSum int) bool {

// 	if root == nil {
// 		return false
// 	}

// 	var dfs func(*TreeNode, int) bool
// 	dfs = func(node *TreeNode, count int) bool {
// 		if node.Left == nil && node.Right == nil {
// 			if count == 0 {
// 				return true
// 			}
// 			return false
// 		}

// 		if node.Left != nil {

// 			if dfs(node.Left, count-node.Left.Val) {
// 				return true
// 			}
// 		}

// 		if node.Right != nil {
// 			if dfs(node.Right, count-node.Right.Val) {
// 				return true
// 			}
// 		}

// 		return false
// 	}

// 	return dfs(root, targetSum-root.Val)
// }

// BFS
// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	queue := list.New()
// 	queue.PushBack(root)
// 	for queue.Len() > 0 {

// 		item := queue.Front()
// 		queue.Remove(item)
// 		node := item.Value.(*TreeNode)
// 		if node.Left == nil && node.Right == nil {
// 			if node.Val == targetSum {
// 				return true
// 			}
// 		}

// 		if node.Left != nil {
// 			node.Left.Val += node.Val
// 			queue.PushBack(node.Left)
// 		}
// 		if node.Right != nil {
// 			node.Right.Val += node.Val
// 			queue.PushBack(node.Right)
// 		}

// 	}
// 	return false

// }

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var bfs func(*TreeNode, int) bool
	bfs = func(n *TreeNode, sum int) bool {
		sum += n.Val
		if sum == targetSum && n.Left == nil && n.Right == nil {
			return true
		}
		if n.Left != nil {
			if bfs(n.Left, sum) {
				return true
			}

		}
		if n.Right != nil {
			if bfs(n.Right, sum) {
				return true
			}
		}
		return false
	}
	return bfs(root, 0)
}

// @lc code=end
