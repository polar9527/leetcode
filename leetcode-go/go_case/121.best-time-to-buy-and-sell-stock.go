package go_case

import "math"

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (53.56%)
 * Likes:    30192
 * Dislikes: 1101
 * Total Accepted:    4.5M
 * Total Submissions: 8.3M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * You are given an array prices where prices[i] is the price of a given stock
 * on the i^th day.
 *
 * You want to maximize your profit by choosing a single day to buy one stock
 * and choosing a different day in the future to sell that stock.
 *
 * Return the maximum profit you can achieve from this transaction. If you
 * cannot achieve any profit, return 0.
 *
 *
 * Example 1:
 *
 *
 * Input: prices = [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit
 * = 6-1 = 5.
 * Note that buying on day 2 and selling on day 1 is not allowed because you
 * must buy before you sell.
 *
 *
 * Example 2:
 *
 *
 * Input: prices = [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transactions are done and the max profit =
 * 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= prices.length <= 10^5
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

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	greedy := func(prices []int) int {
		low := math.MaxInt32
		result := 0
		for i := 0; i < len(prices); i++ {
			low = min(low, prices[i])
			result = max(result, prices[i]-low)
		}

		return result
		// min := prices[0]
		// res := 0
		// for i := 1; i < len(prices); i++ {
		// 	if prices[i]-min > res {
		// 		res = prices[i] - min
		// 	}
		// 	if min > prices[i] {
		// 		min = prices[i]
		// 	}
		// }
		// return res
	}

	dp := func(prices []int) int {
		dp := make([][2]int, len(prices))
		// 持有股票
		dp[0][0] = -prices[0]
		// 不持有股票
		dp[0][1] = 0
		for i := 1; i < len(prices); i++ {
			dp[i][0] = max(dp[i-1][0], -prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		}

		return dp[len(prices)-1][1]
	}
	d := false
	if d {
		return dp(prices)
	}
	return greedy(prices)
}

// @lc code=end
