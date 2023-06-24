package app

import "testing"

func TestM(t *testing.T) {
	var tests = []struct {
		nums   []int
		answer int
	}{
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
		{
			[]int{1, 1, 2},
			2,
		},
	}

	for _, test := range tests {
		result := M(test.nums)
		if result != test.answer {
			t.Errorf("nums = %v, answer = %d, result = %d", test.nums, test.answer, result)
		}
	}
}

func BenchmarkM(b *testing.B) {
	var tests = []struct {
		nums   []int
		answer int
	}{
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
		{
			[]int{1, 1, 2},
			2,
		},
	}

	for i := 0; i < b.N; i++ {
		M(tests[0].nums)
	}
}

func BenchmarkO1(b *testing.B) {
	var tests = []struct {
		nums   []int
		answer int
	}{
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
		{
			[]int{1, 1, 2},
			2,
		},
	}

	for i := 0; i < b.N; i++ {
		O1(tests[0].nums)
	}
}

func BenchmarkO2(b *testing.B) {
	var tests = []struct {
		nums   []int
		answer int
	}{
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
		{
			[]int{1, 1, 2},
			2,
		},
	}

	for i := 0; i < b.N; i++ {
		O2(tests[0].nums)
	}
}
