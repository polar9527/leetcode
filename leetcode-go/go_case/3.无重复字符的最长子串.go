package go_case

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (39.13%)
 * Likes:    9418
 * Dislikes: 0
 * Total Accepted:    2.4M
 * Total Submissions: 6.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lc code=start
// func lengthOfLongestSubstring(s string) int {
// 	// 哈希集合，记录每个字符是否出现过
// 	m := map[byte]int{}
// 	n := len(s)
// 	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
// 	rk, ans := -1, 0
// 	for i := 0; i < n; i++ {
// 		if i != 0 {
// 			// 左指针向右移动一格，移除一个字符
// 			delete(m, s[i-1])
// 		}
// 		for rk+1 < n && m[s[rk+1]] == 0 {
// 			// 不断地移动右指针
// 			m[s[rk+1]]++
// 			rk++
// 		}
// 		// 第 i 到 rk 个字符是一个极长的无重复字符子串
// 		ans = max(ans, rk-i+1)
// 	}
// 	return ans
// }

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	fast := 0
	charMap := make(map[byte]int, 0)
	res := 0
	for slow := 0; slow < n; slow++ {
		// 保证 [slow, fast) 半开半闭 区间内所有的字符唯一
		if slow > 0 {
			delete(charMap, s[slow-1])
		}
		for fast < n && charMap[s[fast]] == 0 {
			charMap[s[fast]]++
			fast++
		}
		// 此时 如果 fast != n, 那意味着 s[fast]和 s[slow:fast]中的某个字符相同
		if res < fast-slow {
			res = fast - slow
		}
	}
	return res
}

// @lc code=end
