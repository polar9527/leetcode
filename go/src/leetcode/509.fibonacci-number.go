/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 *
 * https://leetcode.com/problems/fibonacci-number/description/
 *
 * algorithms
 * Easy (66.80%)
 * Likes:    470
 * Dislikes: 180
 * Total Accepted:    160.8K
 * Total Submissions: 240.3K
 * Testcase Example:  '2'
 *
 * The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
 * Fibonacci sequence, such that each number is the sum of the two preceding
 * ones, starting from 0 and 1. That is,
 *
 *
 * F(0) = 0,   F(1) = 1
 * F(N) = F(N - 1) + F(N - 2), for N > 1.
 *
 *
 * Given N, calculate F(N).
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 1
 * Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 2
 * Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
 *
 *
 * Example 3:
 *
 *
 * Input: 4
 * Output: 3
 * Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
 *
 *
 *
 *
 * Note:
 *
 * 0 ≤ N ≤ 30.
 *
 */
package main

// @lc code=start
func fib(n int) int {
	if n < 0 {
		return -1
	}
	result := []int{0, 1}
	if n < 2 {
		return result[n]
	}

	fibN := 0
	fib1 := 0
	fib2 := 1
	for i := 2; i <= n; i++ {
		fibN = fib2 + fib1
		fib1 = fib2
		fib2 = fibN
	}
	return fibN
}

// @lc code=end
