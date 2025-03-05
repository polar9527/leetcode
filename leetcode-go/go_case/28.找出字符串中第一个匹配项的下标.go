/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 *
 * https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
 *
 * algorithms
 * Easy (43.81%)
 * Likes:    2369
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.8M
 * Testcase Example:  '"sadbutsad"\n"sad"'
 *
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0
 * 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：haystack = "sadbutsad", needle = "sad"
 * 输出：0
 * 解释："sad" 在下标 0 和 6 处匹配。
 * 第一个匹配项的下标是 0 ，所以返回 0 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：haystack = "leetcode", needle = "leeto"
 * 输出：-1
 * 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack 和 needle 仅由小写英文字符组成
 *
 *
 */

// @lc code=start
// brutal, O(n*m)
// func strStr(haystack string, needle string) int {
// 	s, p := len(haystack), len(needle)
// loop:
// 	for i := 0; i < s-p+1; i++ {
// 		for j := 0; j < p; j++ {
// 			if haystack[i+j] != needle[j] {
// 				continue loop
// 			}
// 		}
// 		return i
// 	}
// 	return -1
// }

// KMP 算法， O(m+n)
// next[i] 当前 位置  最长公共前缀后缀长
func strStr(haystack string, needle string) int {
	// prefix array
	getNext := func(next []int, pattern string) {
		// inti
		// j 前缀尾
		j := 0
		next[0] = j
		// i 后追尾
		for i := 1; i < len(pattern); i++ {
			for j > 0 && pattern[i] != pattern[j] {
				j = next[j-1]
			}
			if pattern[i] == pattern[j] {
				j++
			}
			next[i] = j
		}
	}

	n := len(needle)
	if n == 0 {
		return 0
	}

	j := 0
	next := make([]int, n)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

// @lc code=end

