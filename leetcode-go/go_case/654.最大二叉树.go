package go_case

import "math"

/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
 *
 * https://leetcode.cn/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (82.06%)
 * Likes:    783
 * Dislikes: 0
 * Total Accepted:    253.8K
 * Total Submissions: 309.3K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * 给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
 *
 *
 * 创建一个根节点，其值为 nums 中的最大值。
 * 递归地在最大值 左边 的 子数组前缀上 构建左子树。
 * 递归地在最大值 右边 的 子数组后缀上 构建右子树。
 *
 *
 * 返回 nums 构建的 最大二叉树 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,2,1,6,0,5]
 * 输出：[6,3,5,null,2,0,null,null,1]
 * 解释：递归调用如下所示：
 * - [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
 * ⁠   - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
 * ⁠       - 空数组，无子节点。
 * ⁠       - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
 * ⁠           - 空数组，无子节点。
 * ⁠           - 只有一个元素，所以子节点是一个值为 1 的节点。
 * ⁠   - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
 * ⁠       - 只有一个元素，所以子节点是一个值为 0 的节点。
 * ⁠       - 空数组，无子节点。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1]
 * 输出：[3,null,2,null,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * 0 <= nums[i] <= 1000
 * nums 中的所有整数 互不相同
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

// func constructMaximumBinaryTree(nums []int) *TreeNode {
//     if len(nums) == 0 {
//         return nil
//     }
//     // 找到最大值
//     index := findMax(nums)
//     // 构造二叉树
//     root := &TreeNode {
//         Val: nums[index],
//         Left: constructMaximumBinaryTree(nums[:index]),   //左半边
//         Right: constructMaximumBinaryTree(nums[index+1:]),//右半边
//         }
//     return root
// }
// func findMax(nums []int) (index int) {
//     for i, v := range nums {
//         if nums[index] < v {
//             index = i
//         }
//     }
//     return
// }

func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	maxNum := math.MinInt
	index := 0
	for i, n := range nums {
		if n > maxNum {
			maxNum = n
			index = i
		}
	}
	var l, r *TreeNode
	l = constructMaximumBinaryTree(nums[:index])
	r = constructMaximumBinaryTree(nums[index+1:])

	return &TreeNode{
		Val:   nums[index],
		Left:  l,
		Right: r,
	}
}

// @lc code=end
