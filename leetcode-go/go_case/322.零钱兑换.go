/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode.cn/problems/coin-change/description/
 *
 * algorithms
 * Medium (50.33%)
 * Likes:    3022
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 *
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 *
 * 你可以认为每种硬币的数量是无限的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3
 * 解释：11 = 5 + 5 + 1
 *
 * 示例 2：
 *
 *
 * 输入：coins = [2], amount = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：coins = [1], amount = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 *
 *
 */

// @lc code=start
// 1D
// func coinChange(coins []int, amount int) int {

// 	// dp[i][j] 在 [0，i] 硬币中 任意取硬币，凑成 金额 j 所需要的最少 硬币个数
// 	//
// 	// dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i]]+1)
// 	min := func(x, y int) int {
// 		if x < y {
// 			return x
// 		}
// 		return y
// 	}
// 	dp := make([]int, amount+1)
// 	// 初始化为math.MaxInt32
// 	for j := 1; j <= amount; j++ {
// 		dp[j] = math.MaxInt32
// 	}

// 	dp[0] = 0
// 	for j := 1; j <= amount; j++ {
// 		for i := 0; i < len(coins); i++ {
// 			if j-coins[i] >= 0 {
// 				dp[j] = min(dp[j], dp[j-coins[i]]+1)
// 			}
// 		}
// 	}
// 	if dp[amount] == math.MaxInt32 {
// 		return -1
// 	}
// 	return dp[amount]
// amount = 11
// coins = [1,2,5]
// dp
//      			 -> j  按列遍历
//   i              ---------------------------
//  coins[0] = 1   | 0,1,2,2,3,3,2,3,3,4,4,3 -> [5,5,1]
//  coins[1] = 2   | 0,1,1,2,2,3,2,2,3,3,4,3
//  coins[2] = 5   | 0,1,1,2,2,1,2,2,3,3,2,3
// }

// 2D
func coinChange(coins []int, amount int) int {

	// dp[i][j] 在 [0，i] 硬币中 任意取硬币，凑成 金额 j 所需要的最少 硬币个数
	//
	// dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i]]+1)
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	dp := make([][]int, len(coins))
	for i, _ := range dp {
		dp[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	// init j == 0
	for i := 0; i < len(coins); i++ {
		dp[i][0] = 0
	}
	// init i == 0
	for j := 1; j <= amount; j++ {
		if j%coins[0] == 0 {
			dp[0][j] = j / coins[0]
		} else {
			dp[0][j] = -1
		}
	}
	dp[0][0] = 0

	for j := 1; j <= amount; j++ {
		for i := 1; i < len(coins); i++ {
			// 不用 coins[i]
			c0 := dp[i-1][j]

			// 再用一次 coins[i]
			c1 := -1
			if j-coins[i] >= 0 && dp[i][j-coins[i]] != -1 {
				c1 = dp[i][j-coins[i]] + 1
			}

			if c0 == -1 && c1 != -1 {
				dp[i][j] = c1
				continue
			}
			if c0 != -1 && c1 == -1 {
				dp[i][j] = c0
				continue
			}
			if c0 == -1 && c1 == -1 {
				dp[i][j] = -1
				continue
			}
			dp[i][j] = min(c0, c1)
		}
	}

	return dp[len(coins)-1][amount]
	// amount = 11
	// coins = [1,2,5]
	// dp
	//      			 -> j 按层遍历
	//   i              ---------------------------
	//  coins[0] = 1   | 0,1,2,3,4,5,6,7,8,9,10,11
	//  coins[1] = 2   | 0,1,1,2,2,3,3,4,4,5, 5, 6
	//  coins[2] = 5   | 0,1,1,2,2,1,2,2,3,3, 2, 3
}

// @lc code=end

