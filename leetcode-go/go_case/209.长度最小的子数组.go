package go_case

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode.cn/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (46.62%)
 * Likes:    1782
 * Dislikes: 0
 * Total Accepted:    575.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 target 。
 *
 * 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr]
 * ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：target = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：target = 4, nums = [1,4,4]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
 *
 *
 */

// @lc code=start
// func minSubArrayLen(target int, nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	ans := math.MaxInt32
// 	sums := make([]int, n+1)
// 	// 为了方便计算，令 size = n + 1
// 	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
// 	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
// 	// 以此类推
// 	for i := 1; i <= n; i++ {
// 		sums[i] = sums[i-1] + nums[i-1]
// 	}
// 	for i := 1; i <= n; i++ {
// 		t := target + sums[i-1]
// 		bound := sort.SearchInts(sums, t)

// 		ans = min(ans, bound-(i-1))

// 	}
// 	if ans == math.MaxInt32 {
// 		return 0
// 	}
// 	return ans
// }

func minSubArrayLen(target int, nums []int) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	res := math.MaxInt32
	sum := 0
	i := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subl := j - i + 1
			res = min(res, subl)
			sum -= nums[i]
			i++
		}
	}
	if res > len(nums) {
		return 0
	}
	return res
}

// @lc code=end
