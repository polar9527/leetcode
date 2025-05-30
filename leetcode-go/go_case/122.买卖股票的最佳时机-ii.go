/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Medium (74.85%)
 * Likes:    2688
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 1.8M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
 *
 * 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
 *
 * 返回 你能获得的 最大 利润 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：prices = [7,1,5,3,6,4]
 * 输出：7
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
 * 最大总利润为 4 + 3 = 7 。
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
 * 最大总利润为 4 。
 *
 * 示例 3：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 3 * 10^4
 * 0 <= prices[i] <= 10^4
 *
 *
 */

// @lc code=start
// greedy
// func maxProfit(prices []int) int {
// 	sum := 0
// 	for i := 1; i < len(prices); i++ {
// 		delta := prices[i] - prices[i-1]
// 		if delta > 0 {
// 			sum += delta
// 		}
// 	}
// 	return sum
// }

// dp
// func maxProfit(prices []int) int {
// 	dp := make([][2]int, len(prices))

// 	// 持有股票i   时所得的最多现金
// 	// dp[i][0]
// 	// 不持有股票i 所得的最多现金
// 	// dp[i][1]
// 	dp[0][0] = -prices[0]
// 	dp[0][1] = 0
// 	if len(prices) == 1 {
// 		return 0
// 	}
// 	for i := 1; i < len(prices); i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
// 		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
// 	}
// 	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
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
// 			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
// 		}
// 		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
// 	}

// 	return dfs(n-1, 0)
// }

// DP 二维
// func maxProfit(prices []int) int {
// 	n := len(prices)

// 	dp := make([][2]int, n+1)
// 	dp[0][1] = math.MinInt

// 	for i, p := range prices {
// 		dp[i+1][1] = max(dp[i][1], dp[i][0]-p)
// 		dp[i+1][0] = max(dp[i][0], dp[i][1]+p)
// 	}
// 	return dp[n][0]
// }

// DP 一维
func maxProfit(prices []int) int {

	dp0, dp1 := 0, math.MinInt

	for _, p := range prices {
		dp0, dp1 = max(dp0, dp1+p), max(dp1, dp0-p)
	}
	return dp0
}

// @lc code=end

