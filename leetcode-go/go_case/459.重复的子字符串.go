/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 *
 * https://leetcode.cn/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (51.58%)
 * Likes:    1262
 * Dislikes: 0
 * Total Accepted:    312.9K
 * Total Submissions: 600.2K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abab"
 * 输出: true
 * 解释: 可由子串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "aba"
 * 输出: false
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "abcabcabcabc"
 * 输出: true
 * 解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 1 <= s.length <= 10^4
 * s 由小写英文字母组成
 *
 *
 */

// @lc code=start
// brutal
// func repeatedSubstringPattern(s string) bool {
// 	l := len(s)
// 	// 遍历所有可能的 子串长度
// 	for i := 1; i*2 <= l; i++ {
// 		if l%i == 0 {
// 			match := true
// 			for j := i; j < l; j++ {
// 				if s[j] != s[j-i] {
// 					match = false
// 					break
// 				}
// 			}
// 			if match {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// concatenation
// func repeatedSubstringPattern(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	}
// 	t := s + s
// 	return strings.Contains(t[1:len(t)-1], s)
// }

// KMP
func repeatedSubstringPattern(s string) bool {
	l := len(s)
	j := 0
	next := make([]int, l)
	next[0] = j
	for i := 1; i < l; i++ {
		for j > 0 && s[j] != s[i] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
	if next[l-1] != 0 && l%(l-next[l-1]) == 0 {
		return true
	}
	return false

}

// @lc code=end

