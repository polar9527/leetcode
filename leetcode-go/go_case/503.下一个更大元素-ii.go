package go_case

/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 *
 * https://leetcode.cn/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (66.92%)
 * Likes:    845
 * Dislikes: 0
 * Total Accepted:    207.9K
 * Total Submissions: 310.6K
 * Testcase Example:  '[1,2,1]'
 *
 * 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的
 * 下一个更大元素 。
 *
 * 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
 * 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,1]
 * 输出: [2,-1,2]
 * 解释: 第一个 1 的下一个更大的数是 2；
 * 数字 2 找不到下一个更大的数；
 * 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4,3]
 * 输出: [2,3,4,-1,4]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
// func nextGreaterElements(nums []int) []int {
// 	n := len(nums)
// 	ans := make([]int, n)
// 	for i := range ans {
// 		ans[i] = -1
// 	}
// 	stack := []int{}
// 	for i := 0; i < n*2-1; i++ {
// 		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
// 			ans[stack[len(stack)-1]] = nums[i%n]
// 			stack = stack[:len(stack)-1]
// 		}
// 		stack = append(stack, i%n)
// 	}
// 	return ans
// }

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	// 栈中存放的是 nums 的 index
	stack := []int{}
	stack = append(stack, 0)
	//  左闭右闭
	//  i -> [1, n]
	//  i -> [0,n-1]
	// 再增加 n 个， i -> [0, 2n-1]
	for i := 1; i < 2*n; i++ {
		if nums[i%n] <= nums[stack[len(stack)-1]] {
			stack = append(stack, i%n)
		} else {
			for len(stack) != 0 && nums[i%n] > nums[stack[len(stack)-1]] {
				res[stack[len(stack)-1]] = nums[i%n]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i%n)
		}
	}
	return res
}

// @lc code=end
