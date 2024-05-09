package go_case

/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change II
 *
 * https://leetcode.com/problems/coin-change-ii/description/
 *
 * algorithms
 * Medium (63.80%)
 * Likes:    9203
 * Dislikes: 163
 * Total Accepted:    606.9K
 * Total Submissions: 950.9K
 * Testcase Example:  '5\n[1,2,5]'
 *
 * You are given an integer array coins representing coins of different
 * denominations and an integer amount representing a total amount of money.
 *
 * Return the number of combinations that make up that amount. If that amount
 * of money cannot be made up by any combination of the coins, return 0.
 *
 * You may assume that you have an infinite number of each kind of coin.
 *
 * The answer is guaranteed to fit into a signed 32-bit integer.
 *
 *
 * Example 1:
 *
 *
 * Input: amount = 5, coins = [1,2,5]
 * Output: 4
 * Explanation: there are four ways to make up the amount:
 * 5=5
 * 5=2+2+1
 * 5=2+1+1+1
 * 5=1+1+1+1+1
 *
 *
 * Example 2:
 *
 *
 * Input: amount = 3, coins = [2]
 * Output: 0
 * Explanation: the amount of 3 cannot be made up just with coins of 2.
 *
 *
 * Example 3:
 *
 *
 * Input: amount = 10, coins = [10]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= coins.length <= 300
 * 1 <= coins[i] <= 5000
 * All the values of coins are unique.
 * 0 <= amount <= 5000
 *
 *
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))

	for n, _ := range dp {
		dp[n] = make([]int, amount+1)
		dp[n][0] = 1
	}
	for i := coins[0]; i <= amount; i++ {
		dp[0][i] = dp[0][i-coins[0]]
	}

	// for j := 0; j <= amount; j++ {
	// 	if j%coins[0] == 0 {
	// 		dp[0][j] = 1
	// 	}
	// }

	for i := 1; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			}
		}
	}
	return dp[len(coins)-1][amount]
}

// func change(amount int, coins []int) int {
// 	// 定义dp数组
// 	dp := make([]int, amount+1)
// 	// 初始化,0大小的背包, 当然是不装任何东西了, 就是1种方法
// 	dp[0]  = 1
// 	// 遍历顺序
// 	// 遍历物品
// 	for i := 0 ;i < len(coins);i++ {
// 		// 遍历背包
// 		for j:= coins[i] ; j <= amount ;j++ {
// 			// 推导公式
// 			dp[j] += dp[j-coins[i]]
// 		}
// 	}
// 	return dp[amount]
// }

// @lc code=end
