/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (26.02%)
 * Likes:    3587
 * Dislikes: 631
 * Total Accepted:    386.6K
 * Total Submissions: 1.5M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 *
 *
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 * Note:
 *
 *
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 *
 * Example 2:
 *
 *
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the preceding element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 *
 *
 * Example 3:
 *
 *
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 *
 *
 * Example 4:
 *
 *
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore,
 * it matches "aab".
 *
 *
 * Example 5:
 *
 *
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 *
 *
 */
package leetcode

// @lc code=start
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		// if s is "" return true
		// if s is not "" return false
		return len(s) == 0
	}

	firstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		// next pattern
		ret1 := isMatch(s, p[2:])
		// next char
		ret2 := firstMatch && isMatch(s[1:], p)
		return ret1 || ret2
	} else {
		// 	len(p) >= 2 && p[1] != '*'
		ret := firstMatch && isMatch(s[1:], p[1:])
		return ret
	}
}

// @lc code=end
