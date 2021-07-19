/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 *
 * https://leetcode-cn.com/problems/single-number-iii/description/
 *
 * algorithms
 * Medium (72.61%)
 * Likes:    418
 * Dislikes: 0
 * Total Accepted:    47.2K
 * Total Submissions: 65.1K
 * Testcase Example:  '[1,2,1,3,2,5]'
 *
 * 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
 *
 *
 *
 * 进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,1,3,2,5]
 * 输出：[3,5]
 * 解释：[5, 3] 也是有效的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-1,0]
 * 输出：[-1,0]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[1,0]
 *
 *
 * 提示：
 *
 *
 * 2
 * -2^31
 * 除两个只出现一次的整数外，nums 中的其他数字都出现两次
 *
 *
 */

// @lc code=start
func singleNumber(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	// xor 最低的第一个1 的 bit index
	index := 0
	for xor&1 == 0 {
		index++
		xor >>= 1
	}
	flag := 1 << index

	res := []int{0, 0}
	for _, n := range nums {
		if n&flag == 0 {
			res[0] ^= n
		} else {
			res[1] ^= n
		}
	}
	return res
}

// @lc code=end

