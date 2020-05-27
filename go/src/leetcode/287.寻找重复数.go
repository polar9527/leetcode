/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 *
 * https://leetcode-cn.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (63.95%)
 * Likes:    652
 * Dislikes: 0
 * Total Accepted:    68.6K
 * Total Submissions: 105K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和
 * n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
 *
 * 示例 1:
 *
 * 输入: [1,3,4,2,2]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [3,1,3,4,2]
 * 输出: 3
 *
 *
 * 说明：
 *
 *
 * 不能更改原数组（假设数组是只读的）。
 * 只能使用额外的 O(1) 的空间。
 * 时间复杂度小于 O(n^2) 。
 * 数组中只有一个重复的数字，但它可能不止重复出现一次。
 *
 *
 */
package main

// @lc code=start
func findDuplicate(nums []int) int {
	s, f := 0, 0
	for s, f = nums[s], nums[nums[f]]; s != f; s, f = nums[s], nums[nums[f]] {

	}

	s = 0
	for ; s != f; s, f = nums[s], nums[f] {

	}
	return s
}

// func findDuplicate(nums []int) int {
// 	r := len(nums)
// 	l := 1
// 	ret := -1
//
// 	for l <= r {
// 		mid := (l+r) >> 1
// 		cnt := 0
// 		for i := 0; i < len(nums); i++ {
// 			if nums[i] <= mid {
// 				cnt++
// 			}
// 		}
// 		if cnt <= mid {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 		ret = l
// 	}
// 	return ret
// }
// @lc code=end
