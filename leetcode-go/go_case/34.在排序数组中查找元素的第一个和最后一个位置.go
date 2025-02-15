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
	res := []int{-1, -1}

	if len(nums) == 0 {
		return res
	}
	// extreme condition
	if nums[0] > target || nums[len(nums)-1] < target {
		return res
	}

	// A
	// target = 2
	// nums = [3,3]

	// B
	// target = 4
	// nums = [3,3]

	// C
	// target = 3
	// nums = [3,3]

	// D
	// target = 4
	// nums = [3,5]

	// E
	// target = 4
	// nums = [3,4,5]
	getRightBoard := func(nums []int, target int) int {
		rightBoard := -1
		left, right := 0, len(nums)-1
		for left <= right {
			// 不能 用 left + right, 有溢出风险
			// mid := (left + right) >> 1
			mid := left + (right-left)>>1

			if nums[mid] == target {
				rightBoard = mid
				left = mid + 1
			}
			if nums[mid] > target {
				// rightOrder 可能在 [left,mid-1] 范围内
				right = mid - 1
			} else if nums[mid] < target {
				// nums[mid] < target
				// rightOrder 可能在 [mid+1,right] 范围内,
				left = mid + 1
			}

		}
		// A
		// left = 0, right = -1, rightBoard = -1

		// B
		// left = len(nums), right = 1, rightBoard = -1

		// C
		// left = len(nums), right = 1, rightBoard = 1
		return rightBoard
	}
	getLeftBoard := func(nums []int, target int) int {
		leftBoard := -1
		left, right := 0, len(nums)-1
		for left <= right {
			// 不能 用 left + right, 有溢出风险
			// mid := (left + right) >> 1
			mid := left + (right-left)>>1

			if nums[mid] == target {
				leftBoard = mid
				right = mid - 1
			}
			if nums[mid] > target {
				// leftBoard 可能在 [left,mid-1] 范围内
				right = mid - 1
			} else if nums[mid] < target {
				// nums[mid] < target
				// leftBoard 可能在 [mid+1,right] 范围内,
				left = mid + 1
			}

		}
		return leftBoard
	}

	res[0] = getLeftBoard(nums, target)
	res[1] = getRightBoard(nums, target)

	return res

}

// @lc code=end
