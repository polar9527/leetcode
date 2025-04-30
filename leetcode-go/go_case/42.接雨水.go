package go_case

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode.cn/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (62.99%)
 * Likes:    4594
 * Dislikes: 0
 * Total Accepted:    738.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */

// @lc code=start
// func trap(height []int) int {
// 	return trap2P(height)
// }

// func trap2P(height []int) (ans int) {
// 	left, right := 0, len(height)-1
// 	leftMax, rightMax := 0, 0
// 	for left < right {
// 		leftMax = max(leftMax, height[left])
// 		rightMax = max(rightMax, height[right])
// 		if height[left] < height[right] {
// 			ans += leftMax - height[left]
// 			left++
// 		} else {
// 			ans += rightMax - height[right]
// 			right--
// 		}
// 	}
// 	return
// }

// func trapStack(height []int) (ans int) {
// 	stack := []int{}
// 	stackTail := func() int {
// 		return stack[len(stack)-1]
// 	}
// 	stackPop := func() {
// 		stack = stack[:len(stack)-1]
// 	}

// 	for i, h := range height {
// 		for len(stack) > 0 && h > height[stackTail()] {
// 			topIndex := stackTail()
// 			stackPop()
// 			if len(stack) == 0 {
// 				break
// 			}
// 			leftIndex := stackTail()
// 			w := i - leftIndex - 1
// 			h := min(height[leftIndex], height[i]) - height[topIndex]
// 			ans += h * w
// 		}
// 		stack = append(stack, i)
// 	}
// 	return ans

// }

// func trapDP(height []int) (ans int) {
// 	n := len(height)
// 	if n == 0 {
// 		return
// 	}

// 	leftMax := make([]int, n)
// 	leftMax[0] = height[0]
// 	for i := 1; i < n; i++ {
// 		leftMax[i] = max(leftMax[i-1], height[i])
// 	}

// 	rightMax := make([]int, n)
// 	rightMax[n-1] = height[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		rightMax[i] = max(rightMax[i+1], height[i])
// 	}

// 	for i, h := range height {
// 		ans += min(leftMax[i], rightMax[i]) - h
// 	}
// 	return
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 2 points
// func trap(height []int) int {
// 	n := len(height)
// 	if n == 1 {
// 		return 0
// 	}
// 	lmax := make([]int, n)
// 	rmax := make([]int, n)

// 	lmax[0] = height[0]
// 	for i := 1; i < n; i++ {
// 		lmax[i] = max(height[i], lmax[i-1])
// 	}

// 	rmax[n-1] = height[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		rmax[i] = max(height[i], rmax[i+1])
// 	}
// 	sum := 0
// 	for i := 0; i < n; i++ {
// 		count := min(lmax[i], rmax[i]) - height[i]
// 		sum += count
// 	}
// 	return sum
// }

// 2 points
func trap(height []int) int {
	l, r := 0, len(height)-1
	lm := 0
	rm := 0
	res := 0
	for l <= r {
		lm = max(lm, height[l])
		rm = max(rm, height[r])
		if lm < rm {
			res += lm - height[l]
			l++
		} else {
			res += rm - height[r]
			r--
		}
	}
	return res
}

// monotony stack
// func trap(height []int) int {
// 	n := len(height)
// 	if n <= 2 {
// 		return 0
// 	}

// 	sum := 0
// 	stack := []int{0}
// 	for i := 1; i < n; i++ {
// 		// 寻找 右边第一个比当前栈顶元素大的元素
// 		if height[stack[len(stack)-1]] > height[i] {
// 			stack = append(stack, i)
// 		} else if height[stack[len(stack)-1]] == height[i] {
// 			// 水坑中某个平面的水是由当前左边最右侧的高度决定的
// 			// 去掉相同高度元素，避免重复计算
// 			stack = stack[:len(stack)-1]
// 			stack = append(stack, i)
// 		} else {
// 			for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
// 				bottom := height[stack[len(stack)-1]]
// 				stack = stack[:len(stack)-1]
// 				if len(stack) > 0 {
// 					h := min(height[stack[len(stack)-1]], height[i]) - bottom
// 					w := i - stack[len(stack)-1] - 1
// 					sum += h * w
// 				}
// 			}
// 			stack = append(stack, i)
// 		}
// 	}
// 	return sum
// }

// @lc code=end
