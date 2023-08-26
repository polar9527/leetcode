package go_case

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode.cn/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.42%)
 * Likes:    1417
 * Dislikes: 0
 * Total Accepted:    775.3K
 * Total Submissions: 2M
 * Testcase Example:  '4'
 *
 * 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
 *
 * 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
 *
 * 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 4
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 8
 * 输出：2
 * 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= x <= 2^31 - 1
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x-1
	ans := 0
	for l <= r {
		n := (l + r) >> 1
		if n*n == x {
			return n
		}
		if n*n > x {
			r = n - 1
		}
		if n*n < x {
			ans = n
			l = n + 1
		}
	}
	return ans
}

// @lc code=end
