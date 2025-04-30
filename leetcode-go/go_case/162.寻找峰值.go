package go_case

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 *
 * https://leetcode.cn/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (49.31%)
 * Likes:    1103
 * Dislikes: 0
 * Total Accepted:    317.8K
 * Total Submissions: 644.6K
 * Testcase Example:  '[1,2,3,1]'
 *
 * 峰值元素是指其值严格大于左右相邻值的元素。
 *
 * 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
 *
 * 你可以假设 nums[-1] = nums[n] = -∞ 。
 *
 * 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,1]
 * 输出：2
 * 解释：3 是峰值元素，你的函数应该返回其索引 2。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,1,3,5,6,4]
 * 输出：1 或 5
 * 解释：你的函数可以返回索引 1，其峰值元素为 2；
 * 或者返回索引 5， 其峰值元素为 6。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * -2^31 <= nums[i] <= 2^31 - 1
 * 对于所有有效的 i 都有 nums[i] != nums[i + 1]
 *
 *
 */

// @lc code=start

// func findPeakElement(nums []int) int {
// 	n := len(nums)

// 	// 辅助函数，输入下标 i，返回 nums[i] 的值
// 	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
// 	get := func(i int) int {
// 		if i == -1 || i == n {
// 			return math.MinInt64
// 		}
// 		return nums[i]
// 	}

// 	left, right := 0, n-1
// 	for {
// 		mid := (left + right) / 2
// 		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
// 			return mid
// 		}
// 		if get(mid) < get(mid+1) {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}
// }

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	// 这样后面的 mid+1 就不会越界
	l, r := 0, n-2

	// [l,r]
	for l <= r {
		mid := l + (r-l)>>1

		if nums[mid] < nums[mid+1] {
			// 而且当 mid + 1 == n-1 时
			// l = n-1, 意味着这就是峰值，循环结束
			// 并以 l 为结果返回
			l = mid + 1
		} else {
			// nums[mid] > nums[mid+1]
			r = mid - 1
		}
	}
	return l
}

// @lc code=end
