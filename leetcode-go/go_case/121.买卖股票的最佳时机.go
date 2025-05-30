/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (58.45%)
 * Likes:    3757
 * Dislikes: 0
 * Total Accepted:    1.7M
 * Total Submissions: 3M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 *
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 *
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[7,1,5,3,6,4]
 * 输出：5
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 *
 *
 * 示例 2：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
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
// greedy
// func maxProfit(prices []int) int {

// 	low := prices[0]
// 	res := 0
// 	for i := 0; i < len(prices); i++ {
// 		// low 过去的最低值
// 		low = min(low, prices[i])
// 		res = max(res, prices[i]-low)
// 	}
// 	return res
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
// 		dp[i][0] = max(dp[i-1][0], -prices[i])
// 		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
// 	}
// 	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
// }

func maxProfit(prices []int) int {
	dp := make([][3]int, len(prices))
	// 第i天 从来就没有持有过股票
	// dp[i][0]
	// 第i天 持有股票   时所得的最多现金
	// dp[i][1]
	// 第i天 卖出过着之前已经卖出了股票 所得的最多现金
	// dp[i][2]
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	if len(prices) == 1 {
		return 0
	}
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
	}
	return dp[len(prices)-1][2]
}

// @lc code=end

