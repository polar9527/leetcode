package go_case

/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于 K 的子数组
 *
 * https://leetcode.cn/problems/subarray-product-less-than-k/description/
 *
 * algorithms
 * Medium (49.65%)
 * Likes:    716
 * Dislikes: 0
 * Total Accepted:    103.5K
 * Total Submissions: 208.5K
 * Testcase Example:  '[10,5,2,6]\n100'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,5,2,6], k = 100
 * 输出：8
 * 解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
 * 需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3], k = 0
 * 输出：0
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * 1 <= nums[i] <= 1000
 * 0 <= k <= 10^6
 *
 *
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return
	}
	prod := 1
	for slow, fast := 0, 0; fast < len(nums); fast++ {
		prod *= nums[fast]
		// 由于 prod 是 [slow:fast] 区间的数之乘积，
		// 所以当 slow == fast 的时候， prod == nums[slow]
		// 继续循环 prod == 1，而 k >= 0，
		// 如果省略 判断条件 slow <= fast
		// 循环会一直进行下去，导致越界
		for prod >= k && slow <= fast {
			prod /= nums[slow]
			slow++
		}
		ans += fast - slow + 1
	}
	return
}

// @lc code=end
