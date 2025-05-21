/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [968] 监控二叉树
 *
 * https://leetcode.cn/problems/binary-tree-cameras/description/
 *
 * algorithms
 * Hard (52.77%)
 * Likes:    774
 * Dislikes: 0
 * Total Accepted:    97.6K
 * Total Submissions: 184.8K
 * Testcase Example:  '[0,0,null,0,0]'
 *
 * 给定一个二叉树，我们在树的节点上安装摄像头。
 *
 * 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
 *
 * 计算监控树的所有节点所需的最小摄像头数量。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[0,0,null,0,0]
 * 输出：1
 * 解释：如图所示，一台摄像头足以监控所有节点。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：[0,0,null,0,null,0,null,null,0]
 * 输出：2
 * 解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。
 *
 *
 *
 * 提示：
 *
 *
 * 给定树的节点数的范围是 [1, 1000]。
 * 每个节点的值都是 0。
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
func minCameraCover(root *TreeNode) int {
	count := 0
	var travel func(*TreeNode) int
	travel = func(n *TreeNode) int {

		if n == nil {
			return 0
		}

		// 0 -> 有覆盖
		// 1 -> 无覆盖
		// 2 -> 有监控

		l := travel(n.Left)
		r := travel(n.Right)

		// TODO
		// 左右节点都有覆盖
		if l == 0 && r == 0 {
			return 1
		}

		// 左右中至少有一个无覆盖
		if l == 1 || r == 1 {
			count++
			return 2
		}

		// 在覆盖情况下，左右节点至少有一个有监控
		if l == 2 || r == 2 {
			return 0
		}
		return -1
	}
	if travel(root) == 1 {
		count++
	}
	return count
}

func dfs(node *TreeNode) (int, int, int) {
	if node == nil {
		return math.MaxInt / 2, 0, 0 // 除 2 防止加法溢出
	}
	lChoose, lByFa, lByChildren := dfs(node.Left)
	rChoose, rByFa, rByChildren := dfs(node.Right)
	choose := min(lChoose, lByFa) + min(rChoose, rByFa) + 1
	byFa := min(lChoose, lByChildren) + min(rChoose, rByChildren)
	byChildren := min(lChoose+rByChildren, lByChildren+rChoose, lChoose+rChoose)
	return choose, byFa, byChildren
}

func minCameraCover(root *TreeNode) int {
	choose, _, byChildren := dfs(root)
	return min(choose, byChildren)
}

// @lc code=end

