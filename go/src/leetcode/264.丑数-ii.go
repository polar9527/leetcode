/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 *
 * https://leetcode-cn.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (51.40%)
 * Likes:    268
 * Dislikes: 0
 * Total Accepted:    24.2K
 * Total Submissions: 47K
 * Testcase Example:  '10'
 *
 * 编写一个程序，找出第 n 个丑数。
 *
 * 丑数就是质因数只包含 2, 3, 5 的正整数。
 *
 * 示例:
 *
 * 输入: n = 10
 * 输出: 12
 * 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
 *
 * 说明:
 *
 *
 * 1 是丑数。
 * n 不超过1690。
 *
 *
 */
package main

// @lc code=start
func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}

	uglyNumbers := make([]int, 0)
	uglyNumbers = append(uglyNumbers, 1)
	minuglyNumbers2Index := 0
	minuglyNumbers3Index := 0
	minuglyNumbers5Index := 0

	nextUglyIndex := 1
	for nextUglyIndex < n {
		min := min3(2*uglyNumbers[minuglyNumbers2Index], 3*uglyNumbers[minuglyNumbers3Index], 5*uglyNumbers[minuglyNumbers5Index])
		uglyNumbers = append(uglyNumbers, min)
		for 2*uglyNumbers[minuglyNumbers2Index] <= min {
			minuglyNumbers2Index++
		}
		for 3*uglyNumbers[minuglyNumbers3Index] <= min {
			minuglyNumbers3Index++
		}
		for 5*uglyNumbers[minuglyNumbers5Index] <= min {
			minuglyNumbers5Index++
		}
		nextUglyIndex++
	}
	return uglyNumbers[nextUglyIndex-1]
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

// @lc code=end
