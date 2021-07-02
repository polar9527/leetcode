/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (39.33%)
 * Likes:    696
 * Dislikes: 0
 * Total Accepted:    58.5K
 * Total Submissions: 144.7K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 *
 *
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 *
 *
 *
 *
 *
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 *
 *
 *
 * 示例:
 *
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 *
 */
package leetcode

// @lc code=start
func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return heights[0]
	}

	stack := make([]int, 0, l)
	area := 0
	for i := 0; i < l; i++ {
		for len(stack) != 0 && heights[stack[len(stack)-1]] > heights[i] {
			height := heights[stack[len(stack)-1]]

			for len(stack) != 0 && height == heights[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}

			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			area = max(area, width*height)
		}
		stack = append(stack, i)
	}
	for len(stack) != 0 {
		height := heights[stack[len(stack)-1]]

		for len(stack) != 0 && height == heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		var width int
		if len(stack) == 0 {
			width = l
		} else {
			width = l - stack[len(stack)-1] - 1
		}
		area = max(area, width*height)
	}
	return area
}

// @lc code=end
//
// func main() {
// 	h := []int{2,1,5,6,2,3}
// 	largestRectangleArea(h)
// }
