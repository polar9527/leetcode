package go_case

/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
 *
 * https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/
 *
 * algorithms
 * Medium (70.10%)
 * Likes:    665
 * Dislikes: 0
 * Total Accepted:    220.8K
 * Total Submissions: 315.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
 *
 *
 * 每条从根节点到叶节点的路径都代表一个数字：
 *
 *
 * 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
 *
 *
 * 计算从根节点到叶节点生成的 所有数字之和 。
 *
 * 叶节点 是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3]
 * 输出：25
 * 解释：
 * 从根到叶子节点路径 1->2 代表数字 12
 * 从根到叶子节点路径 1->3 代表数字 13
 * 因此，数字总和 = 12 + 13 = 25
 *
 * 示例 2：
 *
 *
 * 输入：root = [4,9,0,5,1]
 * 输出：1026
 * 解释：
 * 从根到叶子节点路径 4->9->5 代表数字 495
 * 从根到叶子节点路径 4->9->1 代表数字 491
 * 从根到叶子节点路径 4->0 代表数字 40
 * 因此，数字总和 = 495 + 491 + 40 = 1026
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 1000] 内
 * 0
 * 树的深度不超过 10
 *
 *
 *
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func sumNumbers(root *TreeNode) int {
// 	sum := 0
// 	var dfs func(n *TreeNode, num int)
// 	dfs = func(n *TreeNode, num int) {

// 		num += n.Val
// 		if n.Left != nil {
// 			dfs(n.Left, num*10)
// 		}
// 		if n.Right != nil {
// 			dfs(n.Right, num*10)
// 		}
// 		if n.Left == nil && n.Right == nil {
// 			sum += num
// 			return
// 		}
// 	}
// 	dfs(root, 0)
// 	return sum
// }

// DFS
func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

// BFS
// type pair struct {
//     node *TreeNode
//     num  int
// }

// func sumNumbers(root *TreeNode) (sum int) {
//     if root == nil {
//         return
//     }
//     queue := []pair{{root, root.Val}}
//     for len(queue) > 0 {
//         p := queue[0]
//         queue = queue[1:]
//         left, right, num := p.node.Left, p.node.Right, p.num
//         if left == nil && right == nil {
//             sum += num
//         } else {
//             if left != nil {
//                 queue = append(queue, pair{left, num*10 + left.Val})
//             }
//             if right != nil {
//                 queue = append(queue, pair{right, num*10 + right.Val})
//             }
//         }
//     }
//     return
// }

// @lc code=end
