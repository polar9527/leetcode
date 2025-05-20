/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 买卖股票的最佳时机含冷冻期
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (65.11%)
 * Likes:    1838
 * Dislikes: 0
 * Total Accepted:    370.8K
 * Total Submissions: 569.2K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
 *
 * 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
 *
 *
 * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 *
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: prices = [1,2,3,0,2]
 * 输出: 3
 * 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
 *
 * 示例 2:
 *
 *
 * 输入: prices = [1]
 * 输出: 0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 5000
 * 0 <= prices[i] <= 1000
 *
 *
 */

// @lc code=start
// func maxProfit(prices []int) int {
// 	l := len(prices)
// 	if l == 1 {
// 		return 0
// 	}
// 	// dp[i][0] 持有股票的状态
// 	// dp[i][1] 第i天之前就已经卖出股票, 并且第i天仍然不买入,而保持不持有股票的状态
// 	// dp[i][2] 第i天才卖出股票
// 	// dp[i][3] 因为 i-1 天 当天卖出了股票，而导致 第i天进入冷冻期

// 	// dp[i][0] = dp[i-1][0], dp[i-1][1] - prices[i], dp[i][3] - prices[i]
// 	// dp[i][1] = dp[i-1][1], dp[i][3]
// 	// dp[i][2] = dp[i][0] + prices[i]
// 	// dp[i][3] = dp[i-1][2]
// 	dp := make([][4]int, l)
// 	dp[0][0] = -prices[0]
// 	dp[0][1] = 0
// 	dp[0][2] = 0
// 	dp[0][3] = 0

// 	for i := 1; i < l; i++ {
// 		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
// 		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
// 		dp[i][2] = dp[i-1][0] + prices[i]
// 		dp[i][3] = dp[i-1][2]
// 	}
// 	return max(dp[l-1][3], max(dp[l-1][2], dp[l-1][1]))
// }

// dfs
// func maxProfit(prices []int) int {
// 	n := len(prices)

// 	cache := make([][2]int, n)
// 	for i := range cache {
// 		cache[i] = [2]int{-1, -1}
// 	}

// 	var dfs func(int, int) int
// 	dfs = func(i, hold int) (res int) {
// 		if i < 0 {
// 			if hold == 1 {
// 				return math.MinInt
// 			}
// 			return
// 		}
// 		if cache[i][hold] != -1 {
// 			return cache[i][hold]
// 		}

// 		defer func() {
// 			cache[i][hold] = res
// 		}()

// 		if hold == 1 {
// 			return max(dfs(i-1, 1), dfs(i-2, 0)-prices[i])
// 		}
// 		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
// 	}

// 	return dfs(n-1, 0)
// }

// DP 二维
// func maxProfit(prices []int) int {
// 	n := len(prices)

// 	dp := make([][2]int, n+2)
// 	dp[1][1] = math.MinInt

// 	for i, p := range prices {
// 		dp[i+2][1] = max(dp[i+1][1], dp[i][0]-p)
// 		dp[i+2][0] = max(dp[i+1][0], dp[i+1][1]+p)
// 	}
// 	return dp[n+1][0]
// }

// DP 一维
func maxProfit(prices []int) int {

	pre0, dp0, dp1 := 0, 0, math.MinInt

	for _, p := range prices {
		pre0, dp0, dp1 = dp0, max(dp0, dp1+p), max(dp1, pre0-p)
	}
	return dp0
}

// @lc code=end

