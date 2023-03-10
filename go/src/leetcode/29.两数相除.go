/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 *
 * https://leetcode.cn/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (22.21%)
 * Likes:    1074
 * Dislikes: 0
 * Total Accepted:    200.9K
 * Total Submissions: 904.8K
 * Testcase Example:  '10\n3'
 *
 * 给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。
 *
 * 整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。
 *
 * 返回被除数 dividend 除以除数 divisor 得到的 商 。
 *
 * 注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−2^31,  2^31 − 1] 。本题中，如果商 严格大于 2^31 − 1
 * ，则返回 2^31 − 1 ；如果商 严格小于 -2^31 ，则返回 -2^31。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
 *
 * 示例 2:
 *
 *
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。
 *
 *
 *
 * 提示：
 *
 *
 * -2^31 <= dividend, divisor <= 2^31 - 1
 * divisor != 0
 *
 *
 */
package leetcode

import (
	"math"
)

// @lc code=start

func divide(dividend, divisor int) int {

	// 边界条件判断
	// 注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−2^31,  2^31 − 1] 。本题中，如果商 严格大于 2^31 − 1
	// ，则返回 2^31 − 1 ；如果商 严格小于 -2^31 ，则返回 -2^31。

	// 	dividend     | math.MinInt32 | -1 | 1 |  math.MaxInt32   |
	// divisor 	     |				 |    |   | 			     |
	// ------------------------------|----|---|------------------|
	// math.MinInt32 |       1       | 0  | 0 |      0           |
	// ----------------------------------------------------------|
	// -1 			 |      越界      | 1  |-1 |math.MinInt32 + 1 |
	// ----------------------------------------------------------|
	// math.MaxInt32 |       -1      | 0  | 0 |       1          |

	if dividend == 0 {
		return 0
	}
	// 当 dividend 值位 2^31 − 1，divisor在[−2^31,  2^31 − 1]，quotient 都不会越界
	// 当 dividend 值位 −2^31，有越界的可能
	if dividend == math.MinInt32 {
		// 快速返回结果，节约时间
		if divisor == math.MinInt32 {
			return 1
		}
		// 当 dividend 值位 −2^31，divisor值位 -1 的时候，quotient 为 2^31，会造成 **越界**
		if divisor == -1 {
			return math.MaxInt32
		}
		// 快速返回结果，节约时间
		if divisor == 1 {
			return math.MinInt32
		}
		// 快速返回结果，节约时间
		if divisor == math.MaxInt32 {
			return -1
		}
	} else if dividend == math.MaxInt32 {
		// 快速返回结果，节约时间
		if divisor == math.MinInt32 {
			return 0
		}
		// 快速返回结果，节约时间
		if divisor == -1 {
			return math.MinInt32 + 1
		}
		// 快速返回结果，节约时间
		if divisor == 1 {
			return math.MaxInt32
		}
		// 快速返回结果，节约时间
		if divisor == math.MaxInt32 {
			return 1
		}
	}

	// 数值范围是 [−2^31,  2^31 − 1]
	// 负数的范围可以覆盖 正数的范围，反之，就不行
	// 所以全部转成负数进行计算
	sign := false
	if dividend > 0 {
		dividend = -dividend
		sign = !sign
	}
	if divisor > 0 {
		divisor = -divisor
		sign = !sign
	}

	// 用二分法判断 divisor*quotient <= dividend < divisor * (quotient+1）
	result := 0
	for lowerBound, upperBound := 1, math.MaxInt32; lowerBound <= upperBound; {
		medium := lowerBound + (upperBound-lowerBound)>>1
		if quickAdd(dividend, divisor, medium) {
			result = medium
			if medium == math.MaxInt32 {
				break
			}
			lowerBound = medium + 1
		} else {
			upperBound = medium - 1
		}
	}

	if sign {
		return -result
	}
	return result
}

// dividend, divisor 都为负数
// 判断 divisor * quotient >= dividend
func quickAdd(dividend, divisor, quotient int) bool {

	for result, increment := 0, divisor; quotient != 0; quotient >>= 1 {
		// 如果当前的二进制个位是 1，则需要将
		// 假设 quotient == b‘10101’
		// quotient >>= 1, quotient == b‘1010’, 弹出的个位要与quotient相乘，累加到result,
		// 如果弹出的是 b‘1’，则需要累加，如果弹出的是b‘0’， 则不需要累加
		if quotient&1 == 1 {
			if result < dividend-increment {
				return false
			}
			result += increment
		}
		// 在 quotient >>= 1 前，将 quotient 翻倍
		if quotient != 1 {
			increment += increment
		}
	}
	return true
}

// @lc code=end
