package go_case

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (45.05%)
 * Likes:    2521
 * Dislikes: 0
 * Total Accepted:    355.9K
 * Total Submissions: 789.8K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 * 示例 1:
 *
 * 输入：heights = [2,1,5,6,2,3]
 * 输出：10
 * 解释：最大的矩形为图中红色区域，面积为 10
 *
 *
 * 示例 2：
 *
 * 输入： heights = [2,4]
 * 输出： 4
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */

// @lc code=start

// 单调栈
func largestRectangleArea(heights []int) int {
	max := 0
	// 使用切片实现栈
	stack := make([]int, 0)
	// 数组头部加入0
	heights = append([]int{0}, heights...)
	// 数组尾部加入0
	heights = append(heights, 0)
	// 初始化栈，序号从0开始
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		// 结束循环条件为：当即将入栈元素>top元素，也就是形成非单调递增的趋势
		for heights[stack[len(stack)-1]] > heights[i] {
			// mid 是top
			mid := stack[len(stack)-1]
			// 出栈
			stack = stack[0 : len(stack)-1]
			// left是top的下一位元素，i是将要入栈的元素
			left := stack[len(stack)-1]
			// 高度x宽度
			tmp := heights[mid] * (i - left - 1)
			if tmp > max {
				max = tmp
			}
		}
		stack = append(stack, i)
	}
	return max
}

// 双指针
// func largestRectangleArea(heights []int) int {
// 	n := len(heights)
// 	left, right := make([]int, n), make([]int, n)
// 	mono_stack := []int{}
// 	for i := 0; i < n; i++ {
// 		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
// 			mono_stack = mono_stack[:len(mono_stack)-1]
// 		}
// 		if len(mono_stack) == 0 {
// 			left[i] = -1
// 		} else {
// 			left[i] = mono_stack[len(mono_stack)-1]
// 		}
// 		mono_stack = append(mono_stack, i)
// 	}
// 	mono_stack = []int{}
// 	for i := n - 1; i >= 0; i-- {
// 		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
// 			mono_stack = mono_stack[:len(mono_stack)-1]
// 		}
// 		if len(mono_stack) == 0 {
// 			right[i] = n
// 		} else {
// 			right[i] = mono_stack[len(mono_stack)-1]
// 		}
// 		mono_stack = append(mono_stack, i)
// 	}
// 	ans := 0
// 	for i := 0; i < n; i++ {
// 		ans = max(ans, (right[i]-left[i]-1)*heights[i])
// 	}
// 	return ans
// }

// @lc code=end
