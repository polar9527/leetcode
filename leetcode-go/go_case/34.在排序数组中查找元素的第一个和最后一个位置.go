package go_case

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (43.49%)
 * Likes:    2710
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2.3M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
 *
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 *
 * 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 *
 * 示例 3：
 *
 *
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * nums 是一个非递减数组
 * -10^9 <= target <= 10^9
 *
 *
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	getStartBoard := func(nums []int, target int) int {
		//左闭右闭
		left, right := 0, n-1
		for left <= right {
			// 不能 用 left + right, 有溢出风险
			// mid := (left + right) >> 1
			mid := left + (right-left)>>1
			if nums[mid] < target {
				// 范围缩小到 [mid-1, right]
				left = mid + 1
			} else {
				//  target <= nums[mid]
				// 范围缩小到 [left,mid-1] 范围内,
				right = mid - 1
			}
			// 循环不变量
			// nums[left-1] < target
			// target <= nums[right+1]
			// 不确定区域 [左侧全部 < target]  [left, right] [target <= 右侧全部 ]

		}
		return left
	}

	// 如果 nums 中所有数都小于 target
	// start == n
	// 如果 nums 中所有数都大于 target
	// start == 0+
	start := getStartBoard(nums, target)
	// 1. 当 target 大于所有的范围内 nums[i]时，l最后会增加到 n，此时 n-1
	// 2. 当 target 小于所有范围内的 nums[i]时，l最后会一直不动指向 0，此时 r指向-1
	if start == n || nums[start] != target {
		return []int{-1, -1}
	}

	end := getStartBoard(nums, target+1) - 1

	return []int{start, end}

}

// @lc code=end
