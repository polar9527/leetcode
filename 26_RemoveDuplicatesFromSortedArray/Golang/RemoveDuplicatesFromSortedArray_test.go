package app

import "testing"

func TestAppM(t *testing.T) {
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
		result := AppM(test.nums)
		if result != test.answer {
			t.Errorf("nums = %v, answer = %d, result = %d", test.nums, test.answer, result)
		}
	}
}

func BenchmarkAppM(b *testing.B) {
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
		AppM(tests[0].nums)
	}
}

func BenchmarkAppO1(b *testing.B) {
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
		AppO1(tests[0].nums)
	}
}

func BenchmarkAppO2(b *testing.B) {
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
		AppO2(tests[0].nums)
	}
}
