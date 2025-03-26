package go_case

import "strconv"

/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 *
 * https://leetcode.cn/problems/monotone-increasing-digits/description/
 *
 * algorithms
 * Medium (51.41%)
 * Likes:    511
 * Dislikes: 0
 * Total Accepted:    144.2K
 * Total Submissions: 280.4K
 * Testcase Example:  '10'
 *
 * 当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。
 *
 * 给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = 10
 * 输出: 9
 *
 *
 * 示例 2:
 *
 *
 * 输入: n = 1234
 * 输出: 1234
 *
 *
 * 示例 3:
 *
 *
 * 输入: n = 332
 * 输出: 299
 *
 *
 *
 *
 * 提示:
 *
 *
 * 0 <= n <= 10^9
 *
 *
 */

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	sb := []byte(s)

	record := len(sb)
	for i := len(sb) - 1; i > 0; i-- {
		if sb[i-1] > sb[i] {
			sb[i-1] -= 1
			record = i
		}
	}
	for i := record; i < len(sb); i++ {
		sb[i] = '9'
	}
	res, _ := strconv.Atoi(string(sb))
	return res
}

// @lc code=end
