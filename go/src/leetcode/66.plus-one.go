/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (40.89%)
 * Total Accepted:    368.9K
 * Total Submissions: 902K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 */

package leetcode

func plusOne(digits []int) []int {
	lenght := len(digits)
	carry := (digits[lenght-1] + 1) / 10
	digits[lenght-1] = (digits[lenght-1] + 1) % 10

	for i := lenght - 2; i >= 0; i-- {
		carryNext := (digits[i] + carry) / 10
		digits[i] = (digits[i] + carry) % 10
		carry = carryNext

	}
	ret := digits
	if carry == 1 {
		ret = []int{carry}
		ret = append(ret, digits...)
	}
	return ret
}

// func main() {
// 	s := []int{8, 9, 9, 9}
// 	plusOne(s)
// }
