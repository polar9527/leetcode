package go_case

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Medium (66.20%)
 * Likes:    13240
 * Dislikes: 2675
 * Total Accepted:    1.8M
 * Total Submissions: 2.8M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * You are given an integer array prices where prices[i] is the price of a
 * given stock on the i^th day.
 *
 * On each day, you may decide to buy and/or sell the stock. You can only hold
 * at most one share of the stock at any time. However, you can buy it then
 * immediately sell it on the same day.
 *
 * Find and return the maximum profit you can achieve.
 *
 *
 * Example 1:
 *
 *
 * Input: prices = [7,1,5,3,6,4]
 * Output: 7
 * Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit
 * = 5-1 = 4.
 * Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 =
 * 3.
 * Total profit is 4 + 3 = 7.
 *
 *
 * Example 2:
 *
 *
 * Input: prices = [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit
 * = 5-1 = 4.
 * Total profit is 4.
 *
 *
 * Example 3:
 *
 *
 * Input: prices = [7,6,4,3,1]
 * Output: 0
 * Explanation: There is no way to make a positive profit, so we never buy the
 * stock to achieve the maximum profit of 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= prices.length <= 3 * 10^4
 * 0 <= prices[i] <= 10^4
 *
 *
 */

// @lc code=start
func maxProfit(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	greedy := func(prices []int) int {
		res := 0
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				res += prices[i] - prices[i-1]
			}
		}

		return res
	}

	dp := func(prices []int) int {
		dp := make([][2]int, len(prices))

		dp[0][0] = -prices[0]

		for i := 1; i < len(prices); i++ {
			// 持有股票
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
			// 不持有股票
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		}

		return dp[len(prices)-1][1]
	}

	d := true
	if d {
		return dp(prices)
	}
	return greedy(prices)
}

// @lc code=end
