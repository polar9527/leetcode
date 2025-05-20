/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/
 *
 * algorithms
 * Hard (52.30%)
 * Likes:    1261
 * Dislikes: 0
 * Total Accepted:    316.4K
 * Total Submissions: 602.9K
 * Testcase Example:  '2\n[2,4,1]'
 *
 * 给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：k = 2, prices = [2,4,1]
 * 输出：2
 * 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
 *
 * 示例 2：
 *
 *
 * 输入：k = 2, prices = [3,2,6,5,0,3]
 * 输出：7
 * 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
 * ⁠    随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 =
 * 3 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= 100
 * 1 <= prices.length <= 1000
 * 0 <= prices[i] <= 1000
 *
 *
 */

// @lc code=start
// func maxProfit(k int, prices []int) int {

// 	l := len(prices)
// 	if l == 1 || k == 0 {
// 		return 0
// 	}

// 	dp := make([][]int, l)
// 	status := make([]int, (2*k+1)*len(prices))
// 	for i := range dp {
// 		dp[i] = status[:2*k+1]
// 		status = status[2*k+1:]
// 	}

// 	for j := 1; j < 2*k; j += 2 {
// 		dp[0][j] = -prices[0]
// 	}

// 	for i := 1; i < l; i++ {
// 		for j := 1; j < 2*k+1; j += 2 {
// 			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
// 			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]+prices[i])
// 		}
// 	}
// 	return dp[l-1][2*k]
// }

// dfs
func maxProfit(k int, prices []int) int {

	n := len(prices)
	cache := make([][][2]int, n)
	for i := range cache {
		cache[i] = make([][2]int, k+1)
		for j := range cache[i] {
			cache[i][j] = [2]int{-1, -1} // -1 表示还没有计算过
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, hold int) (res int) {
		if j < 0 {
			return math.MinInt
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt
			}
			return
		}

		if cache[i][j][hold] != -1 { // 之前计算过
			return cache[i][j][hold]
		}
		defer func() { cache[i][j][hold] = res }() // 记忆化

		if hold == 1 {
			return max(dfs(i-1, j, 1), dfs(i-1, j, 0)-prices[i])
		}
		return max(dfs(i-1, j, 0), dfs(i-1, j-1, 1)+prices[i])
	}
	return dfs(n-1, k, 0)
}

// @lc code=end

