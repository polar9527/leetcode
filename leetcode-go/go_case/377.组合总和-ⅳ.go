/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 *
 * https://leetcode.cn/problems/combination-sum-iv/description/
 *
 * algorithms
 * Medium (53.69%)
 * Likes:    1122
 * Dislikes: 0
 * Total Accepted:    260.1K
 * Total Submissions: 484.1K
 * Testcase Example:  '[1,2,3]\n4'
 *
 * 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
 *
 * 题目数据保证答案符合 32 位整数范围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3], target = 4
 * 输出：7
 * 解释：
 * 所有可能的组合为：
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 * 请注意，顺序不同的序列被视作不同的组合。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [9], target = 3
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
 * nums 中的所有元素 互不相同
 * 1
 *
 *
 *
 *
 * 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
 *
 */

// @lc code=start
// func combinationSum4(nums []int, target int) int {

// 	dp := make([]int, target+1)
// 	dp[0] = 1
// 	for j := 0; j <= target; j++ {
// 		for i := 0; i < len(nums); i++ {
// 			if j-nums[i] >= 0 {
// 				dp[j] += dp[j-nums[i]]
// 			}
// 		}
// 	}
// 	return dp[target]
// }

// 2D 排列
func combinationSum4(nums []int, target int) int {

	dp := make([][]int, len(nums))
	for i, _ := range dp {
		dp[i] = make([]int, target+1)
	}

	for i := 0; i < len(nums); i++ {
		dp[i][0] = 1
	}

	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j-nums[i] >= 0 {
				if i == 0 {
					dp[i][j] += dp[len(nums)-1][j-nums[i]]
				}
				if i > 0 {
					dp[i][j] = dp[i-1][j] + dp[len(nums)-1][j-nums[i]]
				}
			} else {
				if i > 0 {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[len(nums)-1][target]
}

// @lc code=end

