package offer2

import "math"

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode.cn/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (68.85%)
 * Likes:    1562
 * Dislikes: 0
 * Total Accepted:    448.9K
 * Total Submissions: 652K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i
 * 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: temperatures = [73,74,75,71,69,72,76,73]
 * 输出: [1,1,4,2,1,1,0,0]
 *
 *
 * 示例 2:
 *
 *
 * 输入: temperatures = [30,40,50,60]
 * 输出: [1,1,1,0]
 *
 *
 * 示例 3:
 *
 *
 * 输入: temperatures = [30,60,90]
 * 输出: [1,1,0]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= temperatures.length <= 10^5
 * 30 <= temperatures[i] <= 100
 *
 *
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	return dailyTemperaturesStack(temperatures)
}

func dailyTemperaturesStack(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}

	for i, _ := range temperatures {
		curT := temperatures[i]
		// curT > temperatures[stack[len(stack)-1] 确保 stack 内 保存着由底向顶 单调递减的 元素
		// 一旦 遇到一个当前遍历到的元素大于 栈顶元素 的情况，说明之前所入栈的元素碰到了NGE(Next Greater Element)
		// 清理掉栈内所有 符合条件的元素，这些被清理的栈内元素都小于 当前所遍历到的元素。
		// 清理的同时做 根据题型 做相应的登记处理
		for len(stack) > 0 && curT > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperaturesBrute(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	next := make([]int, 101)

	for i := 0; i < 101; i++ {
		next[i] = math.MaxInt32
	}

	for i := length - 1; i >= 0; i-- {
		warmerIndex := math.MaxInt32
		for t := temperatures[i] + 1; t <= 100; t++ {
			if next[t] < warmerIndex {
				warmerIndex = next[t]
			}
		}
		if warmerIndex < math.MaxInt32 {
			ans[i] = warmerIndex - i
		}
		next[temperatures[i]] = i
	}
	return ans
}

// @lc code=end
