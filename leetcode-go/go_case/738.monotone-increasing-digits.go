package go_case

import "strconv"

/*
 * @lc app=leetcode id=738 lang=golang
 *
 * [738] Monotone Increasing Digits
 *
 * https://leetcode.com/problems/monotone-increasing-digits/description/
 *
 * algorithms
 * Medium (47.88%)
 * Likes:    1295
 * Dislikes: 104
 * Total Accepted:    51.8K
 * Total Submissions: 108.2K
 * Testcase Example:  '10'
 *
 * An integer has monotone increasing digits if and only if each pair of
 * adjacent digits x and y satisfy x <= y.
 *
 * Given an integer n, return the largest number that is less than or equal to
 * n with monotone increasing digits.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 10
 * Output: 9
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1234
 * Output: 1234
 *
 *
 * Example 3:
 *
 *
 * Input: n = 332
 * Output: 299
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 10^9
 *
 *
 */

// @lc code=start
func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N) //将数字转为字符串，方便使用下标
	ss := []byte(s)      //将字符串转为byte数组，方便更改。
	n := len(ss)
	if n <= 1 {
		return N
	}
	flag := n
	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			ss[i-1] -= 1
			flag = i
		}
	}
	for i := flag; i < n; i++ { //从flag开始，将后面的数字全部置为9
		ss[i] = '9'
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}

// @lc code=end
