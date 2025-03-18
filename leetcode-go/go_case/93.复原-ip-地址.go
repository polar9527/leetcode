package go_case

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 *
 * https://leetcode.cn/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (59.16%)
 * Likes:    1404
 * Dislikes: 0
 * Total Accepted:    404.7K
 * Total Submissions: 684K
 * Testcase Example:  '"25525511135"'
 *
 * 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
 * 和 "192.168@1.1" 是 无效 IP 地址。
 *
 *
 * 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能
 * 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "25525511135"
 * 输出：["255.255.11.135","255.255.111.35"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "0000"
 * 输出：["0.0.0.0"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "101023"
 * 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 20
 * s 仅由数字组成
 *
 *
 */

// @lc code=start
// func restoreIpAddresses(s string) []string {
// 	var res []string
// 	var path []string

// 	isValidNum := func(s string) bool {
// 		num, _ := strconv.Atoi(s)
// 		if num >= 0 && num <= 255 {
// 			return true
// 		}
// 		return false
// 	}

// 	var backtracing func(int)
// 	backtracing = func(startIndex int) {

// 		if len(path) == 4 {
// 			if startIndex == len(s) {
// 				str := strings.Join(path, ".")
// 				res = append(res, str)
// 			}
// 			return
// 		}

// 		for i := startIndex; i < len(s); i++ {
// 			// 0开头的话，就不需要继续循环来遍历 0x 和 0xx 的情况了
// 			if i-startIndex > 0 && s[startIndex] == '0' {
// 				break
// 			}
// 			field := s[startIndex : i+1]
// 			// 出现无效数字，这一层也不需要继续遍历了
// 			if !isValidNum(field) {
// 				break
// 			}

// 			path = append(path, field)
// 			backtracing(i + 1)
// 			path = path[:len(path)-1]
// 		}

// 	}
// 	backtracing(0)
// 	return res

// }

func restoreIpAddresses(s string) []string {
	isValidNum := func(s string) bool {
		num, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if num >= 0 && num <= 255 {
			return true
		}
		return false
	}
	var res []string
	l := len(s)
	path := []string{}
	var bt func(int)
	bt = func(start int) {
		if len(path) == 4 {
			if start == l {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		for i := start; i < l; i++ {
			field := s[start : i+1]
			if len(field) > 1 && field[0] == '0' {
				break
			}
			if isValidNum(field) {
				path = append(path, field)
				bt(i + 1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	bt(0)
	return res
}

// @lc code=end
