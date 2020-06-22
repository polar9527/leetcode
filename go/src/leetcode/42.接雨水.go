/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (51.15%)
 * Likes:    1377
 * Dislikes: 0
 * Total Accepted:    112K
 * Total Submissions: 218.1K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 *
 * 示例:
 *
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 *
 */
package main

// @lc code=start
func trapTwoPoint(height []int) int {
	var ans int
	l := len(height)

	left := 0
	right := l - 1
	leftMax, rightMax := 0, 0
	for left <= right {
		if height[left] < height[right] {
			if leftMax < height[left] {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if rightMax < height[right] {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}

func trapStack(height []int) int {
	var ans int
	l := len(height)
	if l < 2 {
		return ans
	}
	stack := make([]int, 0)

	for cur := 0; cur < l; cur++ {
		for len(stack) != 0 && height[cur] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := cur - stack[len(stack)-1] - 1
			bouncedHeight := min(height[cur], height[stack[len(stack)-1]]) - height[top]
			ans += distance * bouncedHeight
		}
		stack = append(stack, cur)
	}
	return ans
}

func trapDynamic(height []int) int {
	var ans int
	l := len(height)
	if l <= 1 {
		return ans
	}
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	maxLeft[0] = height[0]
	for i := 1; i < l; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	maxRight[l-1] = height[l-1]
	for i := l - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	for i := 0; i < l; i++ {
		ans += min(maxRight[i], maxLeft[i]) - height[i]
	}
	return ans
}

func trapBruteForce(height []int) int {
	var ans int
	for i := range height {
		maxLeft := 0
		maxRight := 0

		for l := i; l >= 0; l-- {
			maxLeft = max(maxLeft, height[l])
		}

		for r := i; r < len(height); r++ {
			maxRight = max(maxRight, height[r])
		}
		ans += min(maxRight, maxLeft) - height[i]
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
