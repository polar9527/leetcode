package go_case

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
	res := make([]int, len(temperatures))
	// 初始化栈顶元素为第一个下标索引0
	stack := []int{0}

	for i := 1; i < len(temperatures); i++ {
		top := stack[len(stack)-1]
		if temperatures[i] < temperatures[top] {
			stack = append(stack, i)
		} else if temperatures[i] == temperatures[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[top] {
				res[top] = i - top
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}

// 代码简化版本
// func dailyTemperatures(num []int) []int {
//     ans := make([]int, len(num))
//     stack := []int{}
//     for i, v := range num {
//         // 栈不空，且当前遍历元素 v 破坏了栈的单调性
//         for len(stack) != 0 && v > num[stack[len(stack)-1]] {
//             // pop
//             top := stack[len(stack)-1]
//             stack = stack[:len(stack)-1]

//             ans[top] = i - top
//         }
//         stack = append(stack, i)
//     }
//     return ans
// }

// @lc code=end
