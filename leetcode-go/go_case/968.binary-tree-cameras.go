package go_case

/*
 * @lc app=leetcode id=968 lang=golang
 *
 * [968] Binary Tree Cameras
 *
 * https://leetcode.com/problems/binary-tree-cameras/description/
 *
 * algorithms
 * Hard (46.55%)
 * Likes:    5238
 * Dislikes: 75
 * Total Accepted:    135.9K
 * Total Submissions: 291.8K
 * Testcase Example:  '[0,0,null,0,0]'
 *
 * You are given the root of a binary tree. We install cameras on the tree
 * nodes where each camera at a node can monitor its parent, itself, and its
 * immediate children.
 *
 * Return the minimum number of cameras needed to monitor all nodes of the
 * tree.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [0,0,null,0,0]
 * Output: 1
 * Explanation: One camera is enough to monitor all nodes if placed as shown.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [0,0,null,0,null,0,null,null,0]
 * Output: 2
 * Explanation: At least two cameras are needed to monitor all nodes of the
 * tree. The above image shows one of the valid configurations of camera
 * placement.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 1000].
 * Node.val == 0
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
	//  节点的状态值：
	//    0 表示无覆盖
	//    1 表示 有摄像头
	//    2 表示有覆盖
	// 后序遍历，根据左右节点的情况,来判读 自己的状态
	res := 0
	var dfs func(*TreeNode) int
	dfs = func(cur *TreeNode) int {

		if cur == nil {
			return 2
		}
		l, r := dfs(cur.Left), dfs(cur.Right)
		if l == 2 && r == 2 {
			return 0
		}
		if l == 0 || r == 0 {
			res++
			return 1
		}

		if l == 1 || r == 1 {
			return 2
		}

		return -1
	}

	if dfs(root) == 0 {
		res++
	}

	return res
}

// @lc code=end
