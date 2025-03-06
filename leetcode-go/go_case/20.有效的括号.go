/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.02%)
 * Likes:    4657
 * Dislikes: 0
 * Total Accepted:    2.1M
 * Total Submissions: 4.7M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 每个右括号都有一个对应的相同类型的左括号。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "()"
 *
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "()[]{}"
 *
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(]"
 *
 * 输出：false
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "([])"
 *
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅由括号 '()[]{}' 组成
 *
 *
 */

// @lc code=start
func isValid(s string) bool {

	stack := make([]rune, 0)

	for _, r := range s {
		if r == '[' || r == '{' || r == '(' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			} else if r == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else if r == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else if r == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

