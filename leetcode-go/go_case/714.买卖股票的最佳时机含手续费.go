/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (77.10%)
 * Likes:    1130
 * Dislikes: 0
 * Total Accepted:    315.9K
 * Total Submissions: 409.4K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * 给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
 *
 * 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
 *
 * 返回获得利润的最大值。
 *
 * 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
 * 输出：8
 * 解释：能够达到的最大利润:
 * 在此处买入 prices[0] = 1
 * 在此处卖出 prices[3] = 8
 * 在此处买入 prices[4] = 4
 * 在此处卖出 prices[5] = 9
 * 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,3,7,5,10,3], fee = 3
 * 输出：6
 *
 *
 *
 *
 * 提示：
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

	// dp[i][0] 持有股票
	// dp[i][1] 不持有股票

	// dp[i][0] = dp[i-1][0], dp[i-1][1] - prices[i] - flee
	// dp[i][1] = dp[i-1][1], dp[i-1][0] + prices[i]

	l := len(prices)
	dp := make([][2]int, l)
	dp[0][0] = -prices[0] - fee
	dp[0][1] = 0
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return max(dp[l-1][0], dp[l-1][1])
}

// @lc code=end

