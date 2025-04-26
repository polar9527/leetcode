z/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 *
 * https://leetcode.cn/problems/minimum-moves-to-equal-array-elements/description/
 *
 * algorithms
 * Medium (61.36%)
 * Likes:    593
 * Dislikes: 0
 * Total Accepted:    90.8K
 * Total Submissions: 147.9K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：3
 * 解释：
 * 只需要3次操作（注意每次操作会增加两个元素的值）：
 * [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,1,1]
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == nums.length
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 答案保证符合 32-bit 整数
 * 
 * 
 */

// @lc code=start
// 逆向思维，n-1 个数 减1， 相当于另一个数加 1
func minMoves(nums []int) int {
	minNum := nums[0]
	for i := range nums {
		if minNum > nums[i] {
			minNum = nums[i]
		}
	}

	res := 0
	for i := range nums{
		res += nums[i] - minNum
	}
	return res
}
// @lc code=end

