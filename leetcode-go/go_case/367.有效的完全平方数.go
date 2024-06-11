package go_case

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode.cn/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (44.82%)
 * Likes:    524
 * Dislikes: 0
 * Total Accepted:    234.4K
 * Total Submissions: 523.1K
 * Testcase Example:  '16'
 *
 * 给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
 *
 * 完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。
 *
 * 不能使用任何内置的库函数，如  sqrt 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = 16
 * 输出：true
 * 解释：返回 true ，因为 4 * 4 = 16 且 4 是一个整数。
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = 14
 * 输出：false
 * 解释：返回 false ，因为 3.742 * 3.742 = 14 但 3.742 不是一个整数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num <= 2^31 - 1
 *
 *
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	l, r := 1, num-1
	// ans := 0
	for l <= r {
		n := l + (r-l)>>1
		if n*n == num {
			return true
		}
		if n*n > num {
			r = n - 1
		}
		if n*n < num {
			// ans = n
			l = n + 1
		}
	}
	return false
}

// @lc code=end
