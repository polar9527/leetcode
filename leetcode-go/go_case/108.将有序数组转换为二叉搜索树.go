package go_case

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (78.55%)
 * Likes:    1521
 * Dislikes: 0
 * Total Accepted:    483.3K
 * Total Submissions: 615.1K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-10,-3,0,5,9]
 * 输出：[0,-3,9,-10,null,5]
 * 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,3]
 * 输出：[3,1]
 * 解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 按 严格递增 顺序排列
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
// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	pivot := len(nums) / 2

// 	root := &TreeNode{
// 		Val: nums[pivot],
// 	}
// 	root.Left = sortedArrayToBST(nums[:pivot])
// 	root.Right = sortedArrayToBST(nums[pivot+1:])
// 	return root
// }

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	l, r := 0, len(nums)
	mid := l + (r-l)>>1
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
	return root
}

// @lc code=end
