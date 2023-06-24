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
// 数学法
func singleNumber1(nums []int) int {

	set := make(map[int]struct{})
	sumOfNums := 0
	for _, n := range nums {
		set[n] = struct{}{}
		sumOfNums += n
	}

	sumOfSet := 0
	for k, _ := range set {
		sumOfSet += k
	}

	return (3*sumOfSet - sumOfNums) / 2

}

// 统计二进制位
func singleNumber2(nums []int) int {
	// 把计算限定到32位
	res := int32(0)
	// 遍历每一个二进制位
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, n := range nums {
			// n 在 64位系统中是 int64
			// 而且是有符号的
			// 所以要通过强制类型转换到32位有符号整数
			total += (int32(n) >> i) & 1
		}
		if (total % 3) > 0 {
			res |= 1 << i
		}
	}
	return int(res)
}

// 卡诺图
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

