/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (39.45%)
 * Likes:    412
 * Dislikes: 0
 * Total Accepted:    87.2K
 * Total Submissions: 221.1K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 如果数组中不存在目标值，返回 [-1, -1]。
 *
 * 示例 1:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 *
 * 示例 2:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 *
 */
package main

// @lc code=start
func searchRange(nums []int, target int) []int {
	var res = []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return res
	}
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mi := (lo + hi) / 2
		if nums[mi] >= target {
			// 临界点 lo==hi==mi, [mi:]都大于等于target, [:mi-1] 都小于 target
			hi = mi - 1
		}
		if nums[mi] < target {
			// 临界点 lo==hi==mi, [:mi]都小于target，[mi+1:]都大于等于target
			lo = mi + 1
		}
	}
	if lo < len(nums) && nums[lo] == target {
		res[0] = lo
	}

	lo = 0
	hi = len(nums) - 1
	for lo <= hi {
		mi := (lo + hi) / 2
		if nums[mi] <= target {
			// 临界点 lo==hi==mi, [:mi] 都小于或者等于target, [mi+1:] 都大于target
			lo = mi + 1
		}
		if nums[mi] > target {
			// 临界点 lo==hi==mi, [:mi-1] 都小于等于target, [mi:] 都大于target
			hi = mi - 1
		}
	}
	if hi >= 0 && nums[hi] == target {
		res[1] = hi
	}

	return res
}

// @lc code=end

// func main(){
// 	nums := []int{5,7,7,8,8,10}
// 	target := 8
// 	searchRange(nums, target)
// }
