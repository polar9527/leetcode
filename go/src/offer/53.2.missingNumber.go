package main

func missingNumber(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	lo, hi := 0, l-1
	for lo <= hi {
		mi := (lo + hi) / 2
		if nums[mi] <= mi {
			lo = mi + 1
		}
		if nums[mi] > mi {
			hi = mi - 1
		}
		// if nums[mi] < mi {
		//
		// }
	}
	return lo

}
