/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (43.47%)
 * Likes:    196
 * Dislikes: 0
 * Total Accepted:    107.7K
 * Total Submissions: 245.1K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */
package main

import (
	"strings"
)

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

// func main() {
// 	s := "."
// 	isPalindrome(s)
// }
