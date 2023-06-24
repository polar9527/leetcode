package sort

import (
	"math"
)

// least significant digit
func radixLSD(array []int) []int {
	l := len(array)
	if l < 2 {
		return array
	}

	max := math.MinInt64
	for i := range array {
		if max < array[i] {
			max = array[i]
		}
	}

	digits := 0
	for ; max > 0; max /= 10 {
		digits++
	}
	buckets := make([][]int, 0)
	for i := 0; i < 10; i++ {
		buckets = append(buckets, []int{})
	}
	var a = array
	for i := 1; i < digits+1; i++ {
		for _, v := range a {
			index := v % int(math.Pow10(i)) / int(math.Pow10(i-1))
			buckets[index] = append(buckets[index], v)
		}
		a = []int{}
		for _, b := range buckets {
			a = append(a, b...)
		}

		for i := 0; i < 10; i++ {
			buckets[i] = []int{}
		}
	}

	return a
}
