package main

import (
	"fmt"
)

// 题目：给定整数a1, a2, ......an, 判断是否可以从中选出若干数，使它们的和恰好为k。
//
// 样例：输入n = 4 a = {1, 2, 4, 7} k = 13
//
// 输出 YES （13 = 2 + 4 + 7）
//
// 如果将上述k改为15 则输出NO

func dfs(nums []int, k int) bool {
	if k == 0 {
		return true
	}
	// k != 0
	if len(nums) == 0 {
		return false
	}
	for i := range nums {
		var next []int
		next = append(next, nums[:i]...)
		next = append(next, nums[i+1:]...)
		if dfs(next, k) {
			return true
		}
		if dfs(next, k-nums[i]) {
			return true
		}
	}
	return false
}

func main() {
	// nums := []int{1, 2, 4, 7}
	k := 13
	fmt.Println(dfs([]int{}, k))
}
