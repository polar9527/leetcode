package go_case

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 *
 * https://leetcode-cn.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (71.89%)
 * Likes:    686
 * Dislikes: 0
 * Total Accepted:    92K
 * Total Submissions: 128K
 * Testcase Example:  '[2,2,3,2]'
 *
 * 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,2,3,2]
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,1,0,1,99]
 * 输出：99
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -2^31
 * nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
 *
 *
 *
 *
 * 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 */

// @lc code=start
// 状态转移，卡诺图
func singleNumber(nums []int) int {

	one, two := 0, 0
	for _, n := range nums {
		one = ^two & (n ^ one)
		// 下面的one是已经转换过状态的的one了
		// 即，不是初始输入的one
		two = ^one & (n ^ two)
	}
	return one
}

// @lc code=end
