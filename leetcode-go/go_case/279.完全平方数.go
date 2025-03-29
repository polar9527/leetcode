/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode.cn/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (67.90%)
 * Likes:    2148
 * Dislikes: 0
 * Total Accepted:    678.2K
 * Total Submissions: 997.7K
 * Testcase Example:  '12'
 *
 * 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
 *
 * 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11
 * 不是。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 12
 * 输出：3
 * 解释：12 = 4 + 4 + 4
 *
 * 示例 2：
 *
 *
 * 输入：n = 13
 * 输出：2
 * 解释：13 = 4 + 9
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 *
 *
 */

// @lc code=start
func numSquares(n int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// 背包大小为n
	// 物品为完全平方数
	// dp[j] 和为j 的完全平方数的最少数量
	// dp[j] = min(dp[j],dp[j-i*i]+1)
	dp := make([]int, n+1)
	// for i, _ := range dp {
	// 	dp[i] = math.MaxInt32
	// }
	dp[0] = 0
	for j := 1; j <= n; j++ { //背包
		dp[j] = 10e4 + 1
		for i := 1; i*i <= j; i++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

// @lc code=end

