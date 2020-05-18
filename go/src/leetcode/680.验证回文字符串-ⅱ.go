/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 *
 * https://leetcode-cn.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (36.57%)
 * Likes:    150
 * Dislikes: 0
 * Total Accepted:    20.4K
 * Total Submissions: 55.5K
 * Testcase Example:  '"aba"'
 *
 * 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
 *
 * 示例 1:
 *
 *
 * 输入: "aba"
 * 输出: True
 *
 *
 * 示例 2:
 *
 *
 * 输入: "abca"
 * 输出: True
 * 解释: 你可以删除c字符。
 *
 *
 * 注意:
 *
 *
 * 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
 *
 *
 */
package main

// @lc code=start
func validPalindrome(s string) bool {

	return validPalindromeCore(s, 0)
}

func validPalindromeCore(s string, count int) bool {
	l := len(s)
	if l <= 1 {
		return true
	}
	if s[0] == s[l-1] {
		return validPalindromeCore(s[1:l-1], count)
	} else if count == 0 {
		count = 1
		return validPalindromeCore(s[1:], count) || validPalindromeCore(s[:l-1], count)
	} else {
		return false
	}
}

// @lc code=end
