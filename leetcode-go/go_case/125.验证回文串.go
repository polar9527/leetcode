package go_case

import "strings"

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode.cn/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.40%)
 * Likes:    665
 * Dislikes: 0
 * Total Accepted:    485.4K
 * Total Submissions: 1M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
 *
 * 字母和数字都属于字母数字字符。
 *
 * 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: s = "A man, a plan, a canal: Panama"
 * 输出：true
 * 解释："amanaplanacanalpanama" 是回文串。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "race a car"
 * 输出：false
 * 解释："raceacar" 不是回文串。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = " "
 * 输出：true
 * 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
 * 由于空字符串正着反着读都一样，所以是回文串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s 仅由可打印的 ASCII 字符组成
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}
	s = strings.ToLower(s)
	sTrimed := strings.Trim(s, " \t")
	l := 0
	r := len(sTrimed) - 1
	for l < r {
		for l < r && !isValid(sTrimed[l]) {
			l++
		}
		for l < r && !isValid(sTrimed[r]) {
			r--
		}
		if sTrimed[l] == sTrimed[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isValid(b byte) bool {
	if ('a' <= b && b <= 'z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}

// @lc code=end
