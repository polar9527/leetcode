package main

func singleNumbers(nums []int) []int {
	result := make([]int, 2)
	var two int
	for i := range nums {
		two ^= nums[i]
	}

	mask := two & (-two)

	for i := range nums {
		if mask&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}
