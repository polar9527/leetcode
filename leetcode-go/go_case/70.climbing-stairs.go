package go_case

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (52.85%)
 * Likes:    21524
 * Dislikes: 794
 * Total Accepted:    3.2M
 * Total Submissions: 6M
 * Testcase Example:  '2'
 *
 * You are climbing a staircase. It takes n steps to reach the top.
 *
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 *
 * Example 2:
 *
 *
 * Input: n = 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 45
 *
 *
 */

// @lc code=start
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	var dp [3]int
	dp[0], dp[1], dp[2] = 0, 1, 2
	sum := 0
	for i := 3; i <= n; i++ {
		sum = dp[1] + dp[2]
		dp[1], dp[2] = dp[2], sum
	}
	return dp[2]
}

// @lc code=end
