package main

import (
	"math"
)

func isStraight(nums []int) bool {
	set := make(map[int]int)
	max := 0
	min := 14
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		max = int(math.Max(float64(max), float64(nums[i])))
		min = int(math.Min(float64(min), float64(nums[i])))
		if _, ok := set[nums[i]]; ok {
			return false
		}
		set[nums[i]] = i
	}
	return max-min < 5
}

func main() {
	isStraight([]int{1, 2, 3, 4, 5})
}
