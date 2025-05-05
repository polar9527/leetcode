package go_case

import "math"

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (45.22%)
 * Likes:    2588
 * Dislikes: 0
 * Total Accepted:    443.2K
 * Total Submissions: 980.2K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
 * 。
 *
 *
 *
 * 注意：
 *
 *
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 解释：整个字符串 s 是最小覆盖子串。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "a", t = "aa"
 * 输出: ""
 * 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 *
 *
 *
 * 提示：
 *
 *
 * ^m == s.length
 * ^n == t.length
 * 1 <= m, n <= 10^5
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
 */

// @lc code=start
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	count := len(t)
	minLen := math.MaxInt32
	start := 0

	for right < len(s) {
		if need[s[right]] > 0 {
			count--
		}
		need[s[right]]--
		right++

		for count == 0 {
			// 说明此时 的窗口 包含 t
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			if need[s[left]] == 0 {
				// 说明 left 指向的字符串 在 t 中
				count++
			}
			need[s[left]]++
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}

// func minWindow(s string, t string) string {
// 	window := make(map[byte]int)
// 	targetMap := make(map[byte]int)
// 	for _, tChar := range t {
// 		targetMap[byte(tChar)]++
// 	}

// 	check := func() bool {
// 		for char, count := range targetMap {
// 			if window[char] < count {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	minLen := math.MaxInt32
// 	start := 0
// 	end := 0
// 	for left, right := 0, 0; right < len(s); right++ {
// 		// window 没有 展开只包含 targeMap 之前，
// 		// 不停地扩张 窗口的 right

// 		// 只有当字符存在于 targetMap 中时，window才扩张，否则没有意义
// 		if targetMap[s[right]] > 0 {
// 			window[s[right]]++
// 		}

// 		// 当 window 扩张到刚好包含 targeMap 时候
// 		// 缩减窗口的 left，直到不能再缩减，即保证 window包含 targeMap

// 		for left <= right && check() {
// 			if minLen > right-left+1 {
// 				minLen = right - left + 1
// 				start = left
// 				end = right + 1
// 			}
// 			// 缩减窗口
// 			window[s[left]]--
// 			left++
// 		}

// 	}
// 	// minLen 可嫩等于 math.MaxInt32， 有越界风险
// 	// return s[start : start+minLen]
// 	return s[start:end]
// }

// @lc code=end
