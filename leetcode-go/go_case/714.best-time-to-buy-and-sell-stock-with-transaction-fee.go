package go_case

/*
 * @lc app=leetcode id=714 lang=golang
 *
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (68.41%)
 * Likes:    7023
 * Dislikes: 202
 * Total Accepted:    351.1K
 * Total Submissions: 513.1K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * You are given an array prices where prices[i] is the price of a given stock
 * on the i^th day, and an integer fee representing a transaction fee.
 *
 * Find the maximum profit you can achieve. You may complete as many
 * transactions as you like, but you need to pay the transaction fee for each
 * transaction.
 *
 * Note:
 *
 *
 * You may not engage in multiple transactions simultaneously (i.e., you must
 * sell the stock before you buy again).
 * The transaction fee is only charged once for each stock purchase and
 * sale.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: prices = [1,3,2,8,4,9], fee = 2
 * Output: 8
 * Explanation: The maximum profit can be achieved by:
 * - Buying at prices[0] = 1
 * - Selling at prices[3] = 8
 * - Buying at prices[4] = 4
 * - Selling at prices[5] = 9
 * The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
 *
 *
 * Example 2:
 *
 *
 * Input: prices = [1,3,7,5,10,3], fee = 3
 * Output: 6
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= prices.length <= 5 * 10^4
 * 1 <= prices[i] < 5 * 10^4
 * 0 <= fee < 5 * 10^4
 *
 *
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))

	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// 持有股票
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		// 不持有股票
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return dp[len(prices)-1][1]
}

// @lc code=end
