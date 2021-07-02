/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 *
 * https://leetcode.com/problems/valid-number/description/
 *
 * algorithms
 * Hard (14.62%)
 * Likes:    640
 * Dislikes: 4393
 * Total Accepted:    157.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '"0"'
 *
 * Validate if a given string can be interpreted as a decimal number.
 *
 * Some examples:
 * "0" => true
 * " 0.1 " => true
 * "abc" => false
 * "1 a" => false
 * "2e10" => true
 * " -90e3   " => true
 * " 1e" => false
 * "e3" => false
 * " 6e-1" => true
 * " 99e2.5 " => false
 * "53.5e93" => true
 * " --6 " => false
 * "-+3" => false
 * "95a54e53" => false
 *
 * Note: It is intended for the problem statement to be ambiguous. You should
 * gather all requirements up front before implementing one. However, here is a
 * list of characters that can be in a valid decimal number:
 *
 *
 * Numbers 0-9
 * Exponent - "e"
 * Positive/negative sign - "+"/"-"
 * Decimal point - "."
 *
 *
 * Of course, the context of these characters also matters in the input.
 *
 * Update (2015-02-10):
 * The signature of the C++ function had been updated. If you still see your
 * function signature accepts a const char * argument, please click the reload
 * button to reset your code definition.
 *
 */
package leetcode

import (
	"strings"
)

// @lc code=start
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	index := 0
	return isNumberCore(s, &index)

}

func isNumberCore(s string, index *int) bool {

	numeric := isInteger(s, index)

	if *index < len(s) && s[*index] == '.' {
		*index++
		numeric = isUnsigned(s, index) || numeric
	}

	if *index < len(s) && (s[*index] == 'e' || s[*index] == 'E') {
		*index++
		numeric = numeric && isInteger(s, index)
	}
	return numeric && len(s) == *index
}

func isInteger(s string, index *int) bool {
	if len(s[*index:]) == 0 {
		return false
	}
	if s[*index] == '+' || s[*index] == '-' {
		*index++
	}
	return isUnsigned(s, index)

}

func isUnsigned(s string, index *int) bool {
	if len(s[*index:]) == 0 {
		return false
	}
	start := *index
	for *index < len(s) {
		if s[*index] < '0' || s[*index] > '9' {
			break
		}
		*index++
	}
	return *index > start
}

// @lc code=end
