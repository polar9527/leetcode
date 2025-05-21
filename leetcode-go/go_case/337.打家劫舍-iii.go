/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 *
 * https://leetcode.cn/problems/house-robber-iii/description/
 *
 * algorithms
 * Medium (62.05%)
 * Likes:    2071
 * Dislikes: 0
 * Total Accepted:    395.4K
 * Total Submissions: 636.9K
 * Testcase Example:  '[3,2,3,null,3,null,1]'
 *
 * 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 node 。
 *
 * 除了 node 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果
 * 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
 *
 * 给定二叉树的 node 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: node = [3,2,3,null,3,null,1]
 * 输出: 7
 * 解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
 *
 * 示例 2:
 *
 *
 *
 *
 * 输入: node = [3,4,5,1,3,null,1]
 * 输出: 9
 * 解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 树的节点数在 [1, 10^4] 范围内
 * 0 <= Node.val <= 10^4
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

// recursion with memory
// func rob(root *TreeNode) int {
// 	memory := make(map[*TreeNode]int)
// 	var recursion func(*TreeNode) int
// 	recursion = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		if v, ok := memory[node]; ok {
// 			return v
// 		}
// 		if node.Right == nil && node.Left == nil {
// 			return node.Val
// 		}

// 		// 偷父节点
// 		val1 := node.Val
// 		if node.Left != nil {
// 			val1 += recursion(node.Left.Left) + recursion(node.Left.Right)
// 		}

// 		if node.Right != nil {
// 			val1 += recursion(node.Right.Left) + recursion(node.Right.Right)
// 		}

// 		// 不偷父节点
// 		val2 := recursion(node.Left) + recursion(node.Right)

// 		res := max(val1, val2)
// 		memory[node] = res
// 		return res
// 	}
// 	return recursion(root)
// }

func rob(root *TreeNode) int {

	var travel func(*TreeNode) [2]int
	travel = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{}
		}
		if node.Left == nil && node.Right == nil {
			return [2]int{0, node.Val}
		}

		left := travel(node.Left)
		right := travel(node.Right)

		val0 := max(left[0], left[1]) + max(right[0], right[1])
		val1 := node.Val + left[0] + right[0]
		return [2]int{val0, val1}
	}
	res := travel(root)

	return max(res[0], res[1])
}

func dfs(node *TreeNode) (rob, notRob int) {
	if node == nil { // 递归边界
		return 0, 0 // 没有节点，怎么选都是 0
	}
	lRob, lNotRob := dfs(node.Left)                  // 递归左子树
	rRob, rNotRob := dfs(node.Right)                 // 递归右子树
	rob = lNotRob + rNotRob + node.Val               // 选
	notRob = max(lRob, lNotRob) + max(rRob, rNotRob) // 不选
	return
}

func rob(root *TreeNode) int {
	return max(dfs(root)) // 根节点选或不选的最大值
}

// @lc code=end

