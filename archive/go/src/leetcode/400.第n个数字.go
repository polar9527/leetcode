/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第N个数字
 *
 * https://leetcode-cn.com/problems/nth-digit/description/
 *
 * algorithms
 * Medium (36.55%)
 * Likes:    102
 * Dislikes: 0
 * Total Accepted:    8.3K
 * Total Submissions: 22.6K
 * Testcase Example:  '3'
 *
 * 在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 个数字。
 *
 * 注意:
 * n 是正数且在32位整数范围内 ( n < 2^31)。
 *
 * 示例 1:
 *
 * 输入:
 * 3
 *
 * 输出:
 * 3
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11
 *
 * 输出:
 * 0
 *
 * 说明:
 * 第11个数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是0，它是10的一部分。
 *
 *
 */
package leetcode

import (
	"math"
)

// @lccode=start
func findNthDigit(n int) int {
	digits := 1
	for {
		nums := digits * integerCount(digits)
		if n < nums {
			return digitAtIndex(n, digits)
		}
		n -= nums
		digits++
	}
	return -1
}

// 1 [0:9] 1 10
// 2 [10:99] 10 90
// 3 [100:999] 100 900 ...
func integerCount(digitLength int) int {
	if digitLength == 1 {
		return 10
	}
	count := int(math.Pow(float64(10), float64(digitLength-1)))
	return 9 * count
}

func digitAtIndex(index, digits int) int {
	// index = 1
	// digits = 2
	head := headNumber(digits)
	// head = 10
	offset := index / digits
	// offset = 1
	number := head + offset
	// number = 11
	bits := index % digits
	//bit = 1
	bitsReverse := digits - bits
	for i := 1; i < bitsReverse; i++ {
		number /= 10
	}
	return number % 10
}

// 1 [0:9] 0
// 2 [10:99] 10
// 3 [100:999] 100  ...
func headNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return int(math.Pow(float64(10), float64(digits-1)))
}

// @lc code=end

func main() {
	findNthDigit(11)
}
