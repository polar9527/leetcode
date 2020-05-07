package main

import (
	"fmt"
)

func countDigitOne(n int) int {
	return core(n)
}

func core(num int) int {
	if num == 0 {
		return 0
	}
	if num > 0 && num < 10 {
		return 1
	}
	// num = 21345

	highBitCount := 0
	for temp := num; temp > 9; temp /= 10 {
		highBitCount += 1
	}
	highBitBase := base10(highBitCount)
	// highBitBase = 10000
	// rest = 21345 - 21345/10000 * 10000 + 1 = 1346
	rest := num % highBitBase
	var highestBitCount int
	if num/highBitBase > 1 {
		highestBitCount = highBitBase
	} else {
		highestBitCount = rest + 1
	}
	lowerBitCount := (num / highBitBase) * highBitCount * base10(highBitCount-1)
	result := core(rest)
	return result + highestBitCount + lowerBitCount
}

func base10(n int) int {
	highBitBase := 1
	for temp := n; temp > 0; temp-- {
		highBitBase *= 10
	}
	return highBitBase
}

func main() {
	ret := countDigitOne(12)
	fmt.Println(ret)
}
