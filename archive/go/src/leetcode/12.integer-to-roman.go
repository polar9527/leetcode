/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 *
 * https://leetcode.com/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (53.14%)
 * Likes:    836
 * Dislikes: 2370
 * Total Accepted:    308.3K
 * Total Submissions: 576.4K
 * Testcase Example:  "3"
 *
 * Roman numerals are represented by seven different symbols: I, V, X, L, C, D
 * and M.
 *
 *
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * For example, two is written as II in Roman numeral, just two one"s added
 * together. Twelve is written as, XII, which is simply X + II. The number
 * twenty seven is written as XXVII, which is XX + V + II.
 *
 * Roman numerals are usually written largest to smallest from left to right.
 * However, the numeral for four is not IIII. Instead, the number four is
 * written as IV. Because the one is before the five we subtract it making
 * four. The same principle applies to the number nine, which is written as IX.
 * There are six instances where subtraction is used:
 *
 *
 * I can be placed before V (5) and X (10) to make 4 and 9.
 * X can be placed before L (50) and C (100) to make 40 and 90.
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 *
 *
 * Given an integer, convert it to a roman numeral. Input is guaranteed to be
 * within the range from 1 to 3999.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: "III"
 *
 * Example 2:
 *
 *
 * Input: 4
 * Output: "IV"
 *
 * Example 3:
 *
 *
 * Input: 9
 * Output: "IX"
 *
 * Example 4:
 *
 *
 * Input: 58
 * Output: "LVIII"
 * Explanation: L = 50, V = 5, III = 3.
 *
 *
 * Example 5:
 *
 *
 * Input: 1994
 * Output: "MCMXCIV"
 * Explanation: M = 1000,    CM = 900, XC = 90 and IV = 4.
 *
 */

// @lc code=start
package leetcode

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	t := []string{"I", "V", "X", "L", "C", "D", "M"}

	idx := 0
	ret := ""
	for num > 0 {
		curr := num % 10
		num /= 10
		if curr <= 3 {
			ret += strings.Repeat(t[idx], curr)
		} else if curr == 4 {
			ret += t[idx+1]
			ret += t[idx]
		} else if curr <= 8 {
			ret += strings.Repeat(t[idx], curr-5)
			ret += t[idx+1]
		} else {
			ret += t[idx+2]
			ret += t[idx]
		}
		idx += 2
	}
	result := make([]rune, len(ret))
	for i, r := range ret {
		result[len(ret)-1-i] = r
	}
	return string(result)
}

// @lc code=end
func main() {
	n := 1994
	ret := intToRoman(n)
	fmt.Println(ret)
}
