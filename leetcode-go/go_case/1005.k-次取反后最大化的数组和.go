package go_case

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 *
 * https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/description/
 *
 * algorithms
 * Easy (51.83%)
 * Likes:    481
 * Dislikes: 0
 * Total Accepted:    206.8K
 * Total Submissions: 398.8K
 * Testcase Example:  '[4,2,3]\n1'
 *
 * 给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
 *
 *
 * 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
 *
 *
 * 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
 *
 * 以这种方式修改数组后，返回数组 可能的最大和 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,2,3], k = 1
 * 输出：5
 * 解释：选择下标 1 ，nums 变为 [4,-2,3] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,-1,0,2], k = 3
 * 输出：6
 * 解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [2,-3,-1,5,-4], k = 2
 * 输出：13
 * 解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * -100 <= nums[i] <= 100
 * 1 <= k <= 10^4
 *
 *
 */

// @lc code=start
// func largestSumAfterKNegations(nums []int, k int) int {
// 	sort.Slice(nums, func(i, j int) bool {
// 		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
// 	})

// 	for i := 0; i < len(nums); i++ {
// 		if k > 0 && nums[i] < 0 {
// 			nums[i] = -nums[i]
// 			k--
// 		}
// 	}

// 	if k%2 == 1 {
// 		nums[len(nums)-1] = -nums[len(nums)-1]
// 	}

// 	var res int
// 	for i := 0; i < len(nums); i++ {
// 		res += nums[i]
// 	}
// 	return res
// }

func largestSumAfterKNegations(nums []int, k int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	// 此处，k == 0 或者 k > 0
	// 如果 k > 0
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	var res int
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}

// @lc code=end
