package main

import (
	"math"
)

func findNthDigit(n int) int {
	digits := 1
	for {
		nums := digits * integerCount(digits)
		if n < nums {
			return digitAtIndex(n, digits)
		}
		n -= nums
		digits++
	}
	return -1
}

// 1 [0:9] 1 10
// 2 [10:99] 10 90
// 3 [100:999] 100 900 ...
func integerCount(digitLength int) int {
	if digitLength == 1 {
		return 10
	}
	count := int(math.Pow(float64(10), float64(digitLength-1)))
	return 9 * count
}

func digitAtIndex(index, digits int) int {
	// index = 3
	// digits = 2
	head := headNumber(digits)
	// head = 10
	offset := index / digits
	// offset = 1
	number := head + offset
	// number = 11
	bits := index % digits
	//bit = 1
	bitsReverse := digits - bits
	for i := 1; i < bitsReverse; i++ {
		number /= 10
	}
	return number % 10
}

// 1 [0:9] 0
// 2 [10:99] 10
// 3 [100:999] 100  ...
func headNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return int(math.Pow(float64(10), float64(digits-1)))
}
