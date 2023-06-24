package main

func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}

	uglyNumbers := make([]int, 0)
	uglyNumbers = append(uglyNumbers, 1)
	minuglyNumbers2Index := 0
	minuglyNumbers3Index := 0
	minuglyNumbers5Index := 0

	nextUglyIndex := 1
	for nextUglyIndex < n {
		min := min(2*uglyNumbers[minuglyNumbers2Index], 3*uglyNumbers[minuglyNumbers3Index], 5*uglyNumbers[minuglyNumbers5Index])
		uglyNumbers = append(uglyNumbers, min)
		for 2*uglyNumbers[minuglyNumbers2Index] <= min {
			minuglyNumbers2Index++
		}
		for 3*uglyNumbers[minuglyNumbers3Index] <= min {
			minuglyNumbers3Index++
		}
		for 5*uglyNumbers[minuglyNumbers5Index] <= min {
			minuglyNumbers5Index++
		}
		nextUglyIndex++
	}
	return uglyNumbers[nextUglyIndex-1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func main() {
	nthUglyNumber(10)
}
