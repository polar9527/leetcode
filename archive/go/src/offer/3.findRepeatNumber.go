package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 1, 0, 2, 5, 3}
	ret := findRepeatNumber(arr)
	fmt.Println(ret)
}

func findRepeatNumber(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	for _, v := range nums {
		if v < 0 || v > len(nums)-1 {
			return -1
		}
	}

	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] != nums[i] {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			} else {
				return nums[i]
			}
		}
	}
	return -1
}
