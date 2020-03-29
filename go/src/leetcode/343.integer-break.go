/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 *
 * https://leetcode.com/problems/integer-break/description/
 *
 * algorithms
 * Medium (49.34%)
 * Likes:    840
 * Dislikes: 206
 * Total Accepted:    98.7K
 * Total Submissions: 199.5K
 * Testcase Example:  '2'
 *
 * Given a positive integer n, break it into the sum of at least two positive
 * integers and maximize the product of those integers. Return the maximum
 * product you can get.
 *
 * Example 1:
 *
 *
 *
 * Input: 2
 * Output: 1
 * Explanation: 2 = 1 + 1, 1 × 1 = 1.
 *
 *
 * Example 2:
 *
 *
 * Input: 10
 * Output: 36
 * Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
 *
 * Note: You may assume that n is not less than 2 and not larger than 58.
 *
 *
 */
package main

// @lc code=start
func integerBreak(n int) int {
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 1
	case n == 3:
		return 2
	default:
	}

	dynamicMap := make(map[int]int)
	dynamicMap[0] = 0
	dynamicMap[1] = 1
	dynamicMap[2] = 2
	dynamicMap[3] = 3
	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			product := dynamicMap[j] * dynamicMap[i-j]
			if max < product {
				max = product
			}
		}
		dynamicMap[i] = max
	}
	return dynamicMap[n]

}

// @lc code=end
