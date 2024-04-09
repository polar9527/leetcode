package go_case

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=1005 lang=golang
 *
 * [1005] Maximize Sum Of Array After K Negations
 *
 * https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/
 *
 * algorithms
 * Easy (51.00%)
 * Likes:    1518
 * Dislikes: 112
 * Total Accepted:    87.4K
 * Total Submissions: 171.4K
 * Testcase Example:  '[4,2,3]\n1'
 *
 * Given an integer array nums and an integer k, modify the array in the
 * following way:
 *
 *
 * choose an index i and replace nums[i] with -nums[i].
 *
 *
 * You should apply this process exactly k times. You may choose the same index
 * i multiple times.
 *
 * Return the largest possible sum of the array after modifying it in this
 * way.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [4,2,3], k = 1
 * Output: 5
 * Explanation: Choose index 1 and nums becomes [4,-2,3].
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,-1,0,2], k = 3
 * Output: 6
 * Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [2,-3,-1,5,-4], k = 2
 * Output: 13
 * Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^4
 * -100 <= nums[i] <= 100
 * 1 <= k <= 10^4
 *
 *
 */

// @lc code=start
func largestSumAfterKNegations(nums []int, K int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if K > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			K--
		}
	}

	if K%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

// @lc code=end
