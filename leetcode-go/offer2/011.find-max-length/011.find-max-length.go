package offer2

import (
	common "github.com/polar9527/leetcode/leetcode-go/internal/common"
)

func findMaxLength(nums []int) (maxLength int) {
	// 要给 前缀和 本身就就是 0 的 nums[0...i] 的 子数组机会
	// 所以 mp[0]要初始化为 -1，这样i-prevIndex 才能得到nums[0...i] 的长度
	mp := map[int]int{0: -1}
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prevIndex, has := mp[counter]; has {
			maxLength = common.Max(maxLength, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return
}
