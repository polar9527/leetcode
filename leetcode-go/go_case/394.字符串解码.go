package go_case

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 *
 * https://leetcode.cn/problems/decode-string/description/
 *
 * algorithms
 * Medium (59.32%)
 * Likes:    1987
 * Dislikes: 0
 * Total Accepted:    452.3K
 * Total Submissions: 754.8K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "3[a]2[bc]"
 * 输出："aaabcbc"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "3[a2[c]]"
 * 输出："accaccacc"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "2[abc]3[cd]ef"
 * 输出："abcabccdcdcdef"
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "abc3[cd]xyz"
 * 输出："abccdcdcdxyz"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 30
 * s 由小写英文字母、数字和方括号 '[]' 组成
 * s 保证是一个 有效 的输入。
 * s 中所有整数的取值范围为 [1, 300]
 *
 *
 */

// @lc code=start
// stack
func decodeString(s string) string {

	stacks := []string{}
	stacki := []int{}
	var multi int
	var res string
	var digit []rune
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digit = append(digit, r)
		} else if r == '[' {
			if len(digit) > 0 {
				multi, _ = strconv.Atoi(string(digit))
			}
			stacks = append(stacks, res)
			stacki = append(stacki, multi)
			digit = []rune{}
			res = ""
			multi = 0
		} else if r == ']' {
			s := stacks[len(stacks)-1]
			stacks = stacks[:len(stacks)-1]
			i := stacki[len(stacki)-1]
			stacki = stacki[:len(stacki)-1]
			res = s + strings.Repeat(res, i)
		} else {
			res = res + string(r)
		}
	}
	return res
}

// @lc code=end
