/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode.cn/problems/add-binary/description/
 *
 * algorithms
 * Easy (53.24%)
 * Likes:    979
 * Dislikes: 0
 * Total Accepted:    308.4K
 * Total Submissions: 579.2K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入:a = "11", b = "1"
 * 输出："100"
 *
 * 示例 2：
 *
 *
 * 输入：a = "1010", b = "1011"
 * 输出："10101"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= a.length, b.length <= 10^4
 * a 和 b 仅由字符 '0' 或 '1' 组成
 * 字符串如果不是 "0" ，就不含前导零
 *
 *
 */
package leetcode

import (
	"strconv"
)

// @lc code=start
func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	al, bl := len(a), len(b)
	n := max(al, bl)

	for i := 0; i < n; i++ {
		if i < al {
			carry += int(a[al-i-1] - byte('0'))
		}
		if i < bl {
			carry += int(b[bl-i-1] - byte('0'))
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
