/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (29.08%)
 * Likes:    1223
 * Dislikes: 2728
 * Total Accepted:    420.7K
 * Total Submissions: 1.4M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 *
 * Example 1:
 *
 *
 * Input: 2.00000, 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: 2.10000, 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 * Note:
 *
 *
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 *
 *
 */
package leetcode

import (
	"fmt"
	"math"
)

// @lc code=start
func myPow(x float64, n int) float64 {
	if math.Abs(x-0.0) < 0.00001 {
		return 0.0
	}
	if n == 0 {
		return 1.0
	} else if n < 0 {
		x = 1.0 / x
		n = -n
	}
	ret := 1.0
	for n > 1 {
		if n&1 != 0 {
			ret = ret * x
			n--
		} else {
			x = x * x
			n /= 2
		}
	}

	ret *= x

	return ret
}

// @lc code=end

func main() {
	base := 3.0
	exponential := 0.0
	retL := math.Pow(base, exponential)
	fmt.Println(retL)
}
