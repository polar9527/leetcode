/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (38.05%)
 * Likes:    574
 * Dislikes: 0
 * Total Accepted:    67.6K
 * Total Submissions: 171.6K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */
package leetcode

import (
	"fmt"
)

// @lc code=start
func maxProduct(nums []int) int {
	// 	dp[i] = max(dp[i-1], dp[i-1]*nums[i])
	if len(nums) == 0 {
		return 0
	}

	maxN, minN, ret := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxT, minT := maxN, minN
		maxN = max(max(nums[i], maxT*nums[i]), minT*nums[i])
		minN = min(min(nums[i], minT*nums[i]), maxT*nums[i])
		ret = max(maxN, ret)
	}
	return ret
}

// @lc code=end

func main() {
	nums := []int{-2, 3, -4}
	ret := maxProduct(nums)
	fmt.Println(ret)
}
