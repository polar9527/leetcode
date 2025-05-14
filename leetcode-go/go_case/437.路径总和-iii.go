package go_case

/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 *
 * https://leetcode.cn/problems/path-sum-iii/description/
 *
 * algorithms
 * Medium (49.85%)
 * Likes:    1682
 * Dislikes: 0
 * Total Accepted:    240.1K
 * Total Submissions: 482.9K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
 *
 * 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
 * 输出：3
 * 解释：和等于 8 的路径有 3 条，如图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：3
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [0,1000]
 * -10^9
 * -1000
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
// //  DFS
// // rootSum 是以当前节点为起始的链路和为 targetSum 的路径个数
// func rootSum(root *TreeNode, targetSum int) (res int) {
// 	if root == nil {
// 		return
// 	}
// 	val := root.Val
// 	if val == targetSum {
// 		res++
// 	}
// 	res += rootSum(root.Left, targetSum-val)
// 	res += rootSum(root.Right, targetSum-val)
// 	return
// }

// // tree 中的 某个节点向下路径和为 targetSum
// func pathSum(root *TreeNode, targetSum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	res := rootSum(root, targetSum)
// 	res += pathSum(root.Left, targetSum)
// 	res += pathSum(root.Right, targetSum)
// 	return res
// }

// prefix sum + dfs
func pathSum(root *TreeNode, targetSum int) int {
	var res int
	// map[sum]count
	prefixSumMap := make(map[int]int)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, prefixSum int) {
		if node == nil {
			return
		}
		prefixSum += node.Val
		res += prefixSumMap[prefixSum-targetSum]
		prefixSumMap[prefixSum]++
		dfs(node.Left, prefixSum)
		dfs(node.Right, prefixSum)
		prefixSumMap[prefixSum]--
	}
	// ​​存在1个和为 0 的空路径​​（即不选任何节点时的默认状态）
	prefixSumMap[0] = 1
	dfs(root, 0)
	return res
}

// @lc code=end
