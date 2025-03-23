/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 *
 * https://leetcode.cn/problems/maximum-subarray/description/
 *
 * algorithms
 * Medium (55.96%)
 * Likes:    6976
 * Dislikes: 0
 * Total Accepted:    2M
 * Total Submissions: 3.7M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 子数组 是数组中的一个连续部分。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 输出：6
 * 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [5,4,-1,7,8]
 * 输出：23
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 *
 */

// @lc code=start
// greedy
// func maxSubArray(nums []int) int {
// 	preSum := 0
// 	res := nums[0]

// 	for i := 0; i < len(nums); i++ {
// 		if preSum+nums[i] < nums[i] {
// 			preSum = nums[i]
// 		} else {
// 			preSum += nums[i]
// 		}
// 		if res < preSum {
// 			res = preSum
// 		}
// 	}
// 	return res
// }

// dp
func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] < nums[i] {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

// @lc code=end

