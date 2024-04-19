package go_case

/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 *
 * https://leetcode.com/problems/fibonacci-number/description/
 *
 * algorithms
 * Easy (70.87%)
 * Likes:    7996
 * Dislikes: 351
 * Total Accepted:    1.7M
 * Total Submissions: 2.4M
 * Testcase Example:  '2'
 *
 * The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
 * Fibonacci sequence, such that each number is the sum of the two preceding
 * ones, starting from 0 and 1. That is,
 *
 *
 * F(0) = 0, F(1) = 1
 * F(n) = F(n - 1) + F(n - 2), for n > 1.
 *
 *
 * Given n, calculate F(n).
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: 1
 * Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 3
 * Output: 2
 * Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
 *
 *
 * Example 3:
 *
 *
 * Input: n = 4
 * Output: 3
 * Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 30
 *
 *
 */

// @lc code=start
func fib(n int) int {
	if n < 2 {
		return n
	}
	var dp [2]int
	dp[0], dp[1] = 0, 1
	sum := 0
	for i := 1; i < n; i++ {
		sum = dp[0] + dp[1]
		dp[0], dp[1] = dp[1], sum
	}
	return sum
}

// @lc code=end
