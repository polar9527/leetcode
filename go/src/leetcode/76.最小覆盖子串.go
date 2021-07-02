/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode-cn.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (37.83%)
 * Likes:    612
 * Dislikes: 0
 * Total Accepted:    59K
 * Total Submissions: 155K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。
 *
 * 示例：
 *
 * 输入: S = "ADOBECODEBANC", T = "ABC"
 * 输出: "BANC"
 *
 * 说明：
 *
 *
 * 如果 S 中不存这样的子串，则返回空字符串 ""。
 * 如果 S 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 */
package leetcode

import (
	"math"
)

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
