/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 *
 * https://leetcode.cn/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (57.48%)
 * Likes:    657
 * Dislikes: 0
 * Total Accepted:    350.4K
 * Total Submissions: 602.8K
 * Testcase Example:  '"abcdefg"\n2'
 *
 * 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
 *
 *
 * 如果剩余字符少于 k 个，则将剩余字符全部反转。
 * 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abcdefg", k = 2
 * 输出："bacdfeg"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcd", k = 2
 * 输出："bacd"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅由小写英文组成
 * 1 <= k <= 10^4
 *
 *
 */

// @lc code=start
func reverseStr(s string, k int) string {
	reverse := func(s []byte) {
		l, r := 0, len(s)-1
		for l < r {
			// s[l], s[r] = s[r], s[l]
			// 构造 原s[l] ^ s[r] 存入当前 s[l]
			s[l] ^= s[r]
			// 当前s[l] ^ 原s[r] 存入 当前s[r]
			// 等价于 原s[l] ^ 原s[r] ^ 原s[r] => 原s[l] 存入 当前s[r]
			// 这时候 当前s[r] == 原s[l]； 且 当前s[l] == 原s[l] ^ 原s[r]

			s[r] ^= s[l]
			// 当前s[r] ^ 当前s[l] 存入 当前s[l]
			// 等价于 原s[l] ^ 原s[l] ^ 原s[r] =>  当前s[l]
			// 得到  原s[r] =>  当前s[l]
			s[l] ^= s[r]

			// 此时互换成功
			l++
			r--
		}
	}
	ss := []byte(s)
	l := len(s)
	for i := 0; i < l; i += (2 * k) {
		if i+k <= l {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:l])
		}
	}
	return string(ss)
}

// @lc code=end

