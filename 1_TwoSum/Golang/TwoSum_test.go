package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}
	for _, test := range tests {
		result := TwoSum(test.nums, test.target)
		if !IntSliceEqualBCE(test.result, result) {
			t.Errorf("nums = %v, target = %d, result = %v", test.nums, test.target, test.result)
		}
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

func BenchmarkTwoSum(b *testing.B) {
	var tests = []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for i := 0; i < b.N; i++ {
		TwoSum(tests[0].nums, tests[0].target)
	}
}
