/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (42.58%)
 * Likes:    361
 * Dislikes: 0
 * Total Accepted:    68.7K
 * Total Submissions: 155.3K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s
 * 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。
 *
 *
 *
 * 示例：
 *
 * 输入：s = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的连续子数组。
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 *
 */
package leetcode

import (
	"math"
	"sort"
)

// @lc code=start

func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	ans := math.MaxInt32
	sum := 0
	head, tail := 0, 0
	for tail < l {
		sum += nums[tail]
		for sum >= s {
			ans = min(ans, tail-head+1)
			sum -= head
			head++
		}
		tail++
	}

	if ans == math.MaxInt32 {
		ans = 0
	}
	return ans
}

func minSubArrayLenNlogN(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	ans := math.MaxInt32
	sums := make([]int, l+1)

	for i := 1; i <= l; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	// n * log(n)
	for i := 1; i <= l; i++ {
		target := s + sums[i-1]
		//  i <= j < l+1  found
		// j == l+1 not found
		j := sort.SearchInts(sums, target)
		if j <= l {
			ans = min(ans, j-i+1)
		}
	}
	if ans == math.MaxInt32 {
		ans = 0
	}
	return ans
}

// @lc code=end
