/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */
package leetcode

// @lc code=start
func singleNumber(nums []int) []int {
	// 异或特性：
	// C = A^B
	// 如果 C 的二进制中某位为1， 那么在A和B中对应的位一定不同
	var xor int
	for i := range nums {
		xor ^= nums[i]
	}

	bitwise := xor
	var indeOfBit uint
	for bitwise != 0 {
		if bitwise&1 != 0 {
			break
		}
		bitwise >>= 1
		indeOfBit++
	}
	conTrastFactor := 1 << indeOfBit

	var arrayL, arrayR []int
	for i := range nums {
		if conTrastFactor&nums[i] == 0 {
			arrayR = append(arrayR, nums[i])
		} else {
			arrayL = append(arrayL, nums[i])
		}
	}

	var result1, result2 = xor, xor
	for i := range arrayL {
		result1 ^= arrayL[i]
	}
	for i := range arrayR {
		result2 ^= arrayR[i]
	}

	return []int{result1, result2}
}

// @lc code=end
