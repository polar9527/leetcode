package go_case

/*
 * @lc app=leetcode id=503 lang=golang
 *
 * [503] Next Greater Element II
 *
 * https://leetcode.com/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (63.83%)
 * Likes:    7834
 * Dislikes: 191
 * Total Accepted:    393.2K
 * Total Submissions: 615.7K
 * Testcase Example:  '[1,2,1]'
 *
 * Given a circular integer array nums (i.e., the next element of
 * nums[nums.length - 1] is nums[0]), return the next greater number for every
 * element in nums.
 *
 * The next greater number of a number x is the first greater number to its
 * traversing-order next in the array, which means you could search circularly
 * to find its next greater number. If it doesn't exist, return -1 for this
 * number.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,1]
 * Output: [2,-1,2]
 * Explanation: The first 1's next greater number is 2;
 * The number 2 can't find next greater number.
 * The second 1's next greater number needs to search circularly, which is also
 * 2.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,3,4,3]
 * Output: [2,3,4,-1,4]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	for i := 0; i < len(result); i++ {
		result[i] = -1
	}
	//单调递减，存储数组下标索引
	stack := make([]int, 0)
	for i := 0; i < length*2; i++ {
		for len(stack) > 0 && nums[i%length] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			result[index] = nums[i%length]
		}
		stack = append(stack, i%length)
	}
	return result
}

// @lc code=end
