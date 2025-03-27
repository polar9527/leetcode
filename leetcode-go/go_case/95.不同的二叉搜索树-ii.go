/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode.cn/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (74.47%)
 * Likes:    1607
 * Dislikes: 0
 * Total Accepted:    204.8K
 * Total Submissions: 274.9K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
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
func generateTrees(n int) []*TreeNode {

	var bt func(int, int) []*TreeNode
	bt = func(l, r int) (res []*TreeNode) {
		// []
		if l > r {
			return []*TreeNode{nil}
		}

		for i := l; i <= r; i++ {

			leftSet := bt(l, i-1)
			rightSet := bt(i+1, r)

			for _, lt := range leftSet {
				for _, rt := range rightSet {
					res = append(res, &TreeNode{
						Val:   i,
						Left:  lt,
						Right: rt,
					})
				}
			}
		}

		return res
	}
	return bt(1, n)
}

// @lc code=end

