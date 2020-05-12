package main

func twoSum(nums []int, target int) []int {
	// 找到最后一个小于 target 的 nums 的序号
	start := 0
	end := len(nums) - 1

	for start <= end {
		mi := (end + start) / 2

		if nums[mi] >= target {
			end = mi - 1
		} else {
			// 	nums[mi] < target
			start = mi + 1
		}
	}
	if end < 0 {
		return nil
	}

	// 从头尾两数子想加逼近
	// reset start
	start = 0
	for start < end {
		if nums[start]+nums[end] < target {
			start++
			continue
		}
		if nums[start]+nums[end] > target {
			end--
			continue
		}
		if nums[start]+nums[end] == target {
			return []int{nums[start], nums[end]}
		}
	}
	return nil
}
