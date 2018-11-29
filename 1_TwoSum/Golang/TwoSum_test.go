package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	if !IntSliceEqualBCE([]int{0, 1}, result) {
		t.Error("TestTwoSum([]int{2, 7, 11, 15}, 9) == false")
	}
}

func IntSliceEqualBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
