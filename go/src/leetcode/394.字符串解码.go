/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 *
 * https://leetcode-cn.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (49.92%)
 * Likes:    289
 * Dislikes: 0
 * Total Accepted:    31K
 * Total Submissions: 61K
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
 * 示例:
 *
 *
 * s = "3[a]2[bc]", 返回 "aaabcbc".
 * s = "3[a2[c]]", 返回 "accaccacc".
 * s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
 *
 *
 */
package main

import (
	"strings"
)

// @lc code=start
type decoder struct {
	str string
	ptr int
}

func (d *decoder) getString() string {
	if d.ptr == len(d.str) || d.str[d.ptr] == ']' {
		return ""
	}

	cur := d.str[d.ptr]
	repTime := 1
	result := ""
	if cur >= '0' && cur <= '9' {
		repTime = d.getDigits()
		// [
		d.ptr++
		singleStr := d.getString()
		// ]
		d.ptr++
		result = strings.Repeat(singleStr, repTime)
	} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') {
		result = string(cur)
		d.ptr++
	}
	return result + d.getString()
}

func (d *decoder) getDigits() int {
	number := 0
	for ; d.str[d.ptr] >= '0' && d.str[d.ptr] <= '9'; d.ptr++ {
		number = number*10 + int(d.str[d.ptr]-'0')
	}
	return number
}

func decodeString(s string) string {
	d := decoder{s, 0}

	return d.getString()
}

// func decodeString(s string) string {
// 	stack := []string{}
// 	var ptr int
//
// 	for ptr < len(s) {
// 		cur := s[ptr]
// 		if cur >= '0' && cur <= '9' {
// 			var num string
// 			for ; s[ptr] >= '0' && s[ptr] <= '9'; ptr++ {
// 				num += string(s[ptr])
// 			}
// 			stack = append(stack, num)
// 		} else if (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') || cur == '[' {
// 			stack = append(stack, string(cur))
// 			ptr++
// 		} else {
// 			ptr++
// 			var sub []string
// 			// pop all letters in []
// 			for stack[len(stack)-1] != "[" {
// 				sub = append(sub, stack[len(stack)-1])
// 				stack = stack[:len(stack)-1]
// 			}
// 			// 出栈后的字符串组需要逆序
// 			for i := 0; i < len(sub)/2; i++ {
// 				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
// 			}
// 			// pop [
// 			stack = stack[:len(stack)-1]
// 			// pop number
// 			repTimes, _ := strconv.Atoi(stack[len(stack)-1])
// 			stack = stack[:len(stack)-1]
// 			stack = append(stack, strings.Repeat(strings.Join(sub, ""), repTimes))
// 		}
//
// 	}
// 	return strings.Join(stack, "")
// }
// @lc code=end

func main() {
	decodeString("3[a2[bc]]")
}
