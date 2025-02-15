package go_case

/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 *
 * https://leetcode.cn/problems/backspace-string-compare/description/
 *
 * algorithms
 * Easy (47.78%)
 * Likes:    737
 * Dislikes: 0
 * Total Accepted:    245.2K
 * Total Submissions: 513.1K
 * Testcase Example:  '"ab#c"\n"ad#c"'
 *
 * 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
 *
 * 注意：如果对空文本输入退格字符，文本继续为空。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ab#c", t = "ad#c"
 * 输出：true
 * 解释：s 和 t 都会变成 "ac"。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "ab##", t = "c#d#"
 * 输出：true
 * 解释：s 和 t 都会变成 ""。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a#c", t = "b"
 * 输出：false
 * 解释：s 会变成 "c"，但 t 仍然是 "b"。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length, t.length <= 200
 * s 和 t 只含有小写字母以及字符 '#'
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你可以用 O(n) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？
 *
 *
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	ls, lt := len(s)-1, len(t)-1
	skips, skipt := 0, 0

	for ls >= 0 || lt >= 0 {
		for ls >= 0 {
			if s[ls] == '#' {
				skips++
				ls--
			} else if skips > 0 {
				skips--
				ls--
			} else {
				break
			}
		}
		for lt >= 0 {
			if t[lt] == '#' {
				skipt++
				lt--
			} else if skipt > 0 {
				skipt--
				lt--
			} else {
				break
			}
		}
		if ls >= 0 && lt >= 0 {
			if s[ls] != t[lt] {
				return false
			}
		} else if ls >= 0 || lt >= 0 {
			return false
		}
		ls--
		lt--

	}
	return true
}

// @lc code=end
