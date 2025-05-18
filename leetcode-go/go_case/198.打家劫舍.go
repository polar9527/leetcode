/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 *
 * https://leetcode.cn/problems/house-robber/description/
 *
 * algorithms
 * Medium (55.65%)
 * Likes:    3219
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,2,3,1]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2：
 *
 *
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */

// @lc code=start
// time O(n)
// space O(n)
// func rob(nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}

// 	if n == 1 {
// 		return nums[0]
// 	}

// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	dp[1] = max(nums[0], nums[1])
// 	for i := 2; i < n; i++ {
// 		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
// 	}
// 	return dp[n-1]
// }

// time O(n)
// space O(1)
// func rob(nums []int) int {
// 	n := len(nums)
// 	dp0 := 0
// 	dp1 := 0
// 	for i := 0; i < n; i++ {

// 		dp0, dp1 = dp1, max(dp1, dp0+nums[i])
// 	}
// 	return dp1
// }

// dfs
func rob(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		res := max(dfs(i-1), dfs(i-2)+nums[i])
		cache[i] = res

		return res
	}
	return dfs(n - 1)
}

// @lc code=end

