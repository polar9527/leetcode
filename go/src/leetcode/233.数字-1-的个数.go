/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 *
 * https://leetcode-cn.com/problems/number-of-digit-one/description/
 *
 * algorithms
 * Hard (35.51%)
 * Likes:    125
 * Dislikes: 0
 * Total Accepted:    7.2K
 * Total Submissions: 20.3K
 * Testcase Example:  '13'
 *
 * 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
 * 
 * 示例:
 * 
 * 输入: 13
 * 输出: 6 
 * 解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。
 * 
 */

// @lc code=start
func countDigitOne(n int) int {
	return core(n)
}

func core(num int) int {
	if num == 0 {
		return 0
	}
	if num > 0 && num < 10 {
		return 1
	}
	// num = 21345

	highBitCount := 0
	for temp := num ; temp > 9 ;temp /= 10 {
		highBitCount += 1
	}
	highBitBase := base10(highBitCount)
	// highBitBase = 10000
	// rest = 21345 - 21345/10000 * 10000 + 1 = 1346
	rest := num%highBitBase
	var highestBitCount int
	if num/highBitBase > 1 {
		highestBitCount = highBitBase
	} else {
		highestBitCount = rest + 1
	}
	lowerBitCount := (num/highBitBase) * highBitCount * base10(highBitCount-1)
	result := core(rest)
	return result + highestBitCount + lowerBitCount
}

func base10(n int) int {
	highBitBase := 1
	for temp := n ; temp > 0 ; temp-- {
		highBitBase *= 10
	}
	return highBitBase
}

// @lc code=end

