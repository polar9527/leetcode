package go_case

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode.cn/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (73.47%)
 * Likes:    1751
 * Dislikes: 0
 * Total Accepted:    377.8K
 * Total Submissions: 514.2K
 * Testcase Example:  '"aab"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aab"
 * 输出：[["a","a","b"],["aa","b"]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a"
 * 输出：[["a"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 16
 * s 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
// func partition(s string) [][]string {
// 	var res [][]string
// 	var path []string
// 	var backtracing func(int)
// 	isPalindrome := func(s string, start, end int) bool {
// 		for start < end {
// 			if s[start] != s[end] {
// 				return false
// 			}
// 			start++
// 			end--
// 		}
// 		return true
// 	}

// 	backtracing = func(startIndex int) {
// 		if startIndex >= len(s) {
// 			res = append(res, append([]string{}, path...))
// 			return
// 		}

// 		for i := startIndex; i < len(s); i++ {
// 			if isPalindrome(s, startIndex, i) {
// 				path = append(path, s[startIndex:i+1])
// 				backtracing(i + 1)
// 				path = path[:len(path)-1]
// 			}
// 		}

// 	}

// 	backtracing(0)
// 	return res
// }

func partition(s string) [][]string {
	n := len(s)
	res := [][]string{}
	// 左闭右闭
	isPal := func(s string, start, end int) bool {
		for start <= end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}
	var bt func(int, []string)
	bt = func(start int, path []string) {
		if start == n {
			res = append(res, append([]string{}, path...))
		}
		for i := start; i < n; i++ {
			if isPal(s, start, i) {
				path = append(path, s[start:i+1])
				bt(i+1, path)
				path = path[:len(path)-1]
			}
		}
	}

	bt(0, []string{})
	return res
}

// @lc code=end
