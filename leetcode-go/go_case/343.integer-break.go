package go_case

/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 *
 * https://leetcode.com/problems/integer-break/description/
 *
 * algorithms
 * Medium (60.29%)
 * Likes:    5052
 * Dislikes: 444
 * Total Accepted:    353.6K
 * Total Submissions: 586.6K
 * Testcase Example:  '2'
 *
 * Given an integer n, break it into the sum of k positive integers, where k >=
 * 2, and maximize the product of those integers.
 *
 * Return the maximum product you can get.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: 1
 * Explanation: 2 = 1 + 1, 1 × 1 = 1.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 10
 * Output: 36
 * Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= n <= 58
 *
 *
 */

// @lc code=start
func integerBreak(n int) int {
	/**
	  动态五部曲
	  1.确定dp下标及其含义
	  2.确定递推公式
	  3.确定dp初始化
	  4.确定遍历顺序
	  5.打印dp
	  **/

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j <= i/2; j++ {
			// i可以差分为i-j和j。由于需要最大值，故需要通过j遍历所有存在的值，取其中最大的值作为当前i的最大值，在求最大值的时候，一个是j与i-j相乘，一个是j与dp[i-j].
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

// @lc code=end
