/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (43.65%)
 * Likes:    1486
 * Dislikes: 52
 * Total Accepted:    308.4K
 * Total Submissions: 683.8K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's
 * sum equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \    / \
 * 7    2  5   1
 *
 *
 * Return:
 *
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
 *
 *
 */
package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	// result := make([][]int, 0)
	path := []int{}
	if root == nil {
		return result
	}

	printPath(&result, &path, sum, 0, root)
	return result
}

func printPath(ret *[][]int, path *[]int, sum, currentSum int, currentNode *TreeNode) {

	currentSum += currentNode.Val
	*path = append(*path, currentNode.Val)
	if currentSum == sum && currentNode.Left == nil && currentNode.Right == nil {
		var temp []int
		for _, v := range *path {
			temp = append(temp, v)
		}
		*ret = append(*ret, temp)
	}

	if currentNode.Left != nil {
		printPath(ret, path, sum, currentSum, currentNode.Left)
	}
	if currentNode.Right != nil {
		printPath(ret, path, sum, currentSum, currentNode.Right)
	}

	*path = (*path)[:len(*path)-1]
}

// @lc code=end
