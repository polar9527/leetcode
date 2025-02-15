package go_case

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode.cn/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (63.53%)
 * Likes:    2388
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 2.2M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [0]
 * 输出: [0]
 *
 *
 *
 * 提示:
 *
 *
 *
 * 1 <= nums.length <= 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你能尽量减少完成的操作次数吗？
 *
 */

// @lc code=start
func moveZeroes(nums []int) {
	l := len(nums)
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for i := slow; i < l; i++ {
		nums[i] = 0
	}
}

// @lc code=end
