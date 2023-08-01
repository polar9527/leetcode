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
	target, window := map[byte]int{}, map[byte]int{}

	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}

	check := func() bool {
		for k, v := range target {
			if window[k] < v {
				return false
			}
		}
		return true
	}

	length := math.MaxInt32
	ansL, ansR := 0, 0
	for l, r := 0, 0; r < len(s); r++ {
		if target[s[r]] > 0 {
			window[s[r]]++
		}
		for check() && l <= r {
			if length > r-l+1 {
				ansL, ansR = l, r+1
				length = r - l + 1
			}
			// shrink left
			window[s[l]] -= 1
			l++
		}
	}

	return s[ansL:ansR]
}

// @lc code=end
